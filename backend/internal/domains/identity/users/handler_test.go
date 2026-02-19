package users

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_Delete(t *testing.T) {
	mockSvc := new(MockUserService)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := NewUserHandler(mockSvc, logger)

	t.Run("Success", func(t *testing.T) {
		r := chi.NewRouter()
		r.Delete("/{userID}", handler.Delete)

		mockSvc.On("DeleteUser", int64(123)).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/123", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		mockSvc.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		r := chi.NewRouter()
		r.Delete("/{userID}", handler.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/abc", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Failed get userID")
	})
}

func TestUserHandler_Get(t *testing.T) {
    mockSvc := new(MockUserService)
    logger := slog.New(slog.NewTextHandler(io.Discard, nil))
    handler := NewUserHandler(mockSvc, logger)

    t.Run("Success", func(t *testing.T) {
        // Setup chi router to parse URL params
        r := chi.NewRouter()
        r.Get("/{userID}", handler.Get)

        mockSvc.On("GetUser", int64(1)).Return(&User{ID: 1}, nil).Once()

        req := httptest.NewRequest(http.MethodGet, "/1", nil)
        rr := httptest.NewRecorder()
        r.ServeHTTP(rr, req)

        assert.Equal(t, http.StatusOK, rr.Code)
        assert.Contains(t, rr.Body.String(), `"id":1`)
    })
}

func TestUserHandler_FindBy(t *testing.T) {
    mockSvc := new(MockUserService)
    handler := NewUserHandler(mockSvc, slog.Default())

    t.Run("Valid Query Params", func(t *testing.T) {
        // Mock expects a filter with limit 5
        mockSvc.On("FindUserBy", mock.MatchedBy(func(f repository.Filter) bool {
            return f.Limit == 5
        })).Return([]User{{ID: 1}}, nil).Once()

        req := httptest.NewRequest(http.MethodGet, "/?limit=5", nil)
        rr := httptest.NewRecorder()
        
        handler.FindBy(rr, req)

        assert.Equal(t, http.StatusOK, rr.Code)
    })
}

func TestUserHandler_Patch(t *testing.T) {
	mockSvc := new(MockUserService)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := NewUserHandler(mockSvc, logger)

	t.Run("Success", func(t *testing.T) {
		r := chi.NewRouter()
		r.Patch("/{userID}", func(w http.ResponseWriter, req *http.Request) {
			// Simulate middleware passing the DTO
			firstName := "Keith"
			dto := PatchUserRequest{
				UserFields: UserFields{FirstName: &firstName},
			}
			handler.Patch(w, req, dto)
		})

		// Expectation: The service receives a User model with ID 1 and the updated name
		mockSvc.On("UpdateUser", mock.MatchedBy(func(u User) bool {
			return u.ID == 1 && *u.FirstName == "Keith"
		})).Return(&User{ID: 1}, nil).Once()

		req := httptest.NewRequest(http.MethodPatch, "/1", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockSvc.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		r := chi.NewRouter()
		r.Patch("/{userID}", func(w http.ResponseWriter, req *http.Request) {
			handler.Patch(w, req, PatchUserRequest{})
		})

		mockSvc.On("UpdateUser", mock.Anything).
			Return(nil, fmt.Errorf("db error")).Once()

		req := httptest.NewRequest(http.MethodPatch, "/1", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Update user failed")
	})
}

func TestUserHandler_Post(t *testing.T) {
    mockSvc := new(MockUserService)
    logger := slog.New(slog.NewTextHandler(io.Discard, nil))
    handler := NewUserHandler(mockSvc, logger)

    t.Run("Created Successfully", func(t *testing.T) {
        username := "keith"
        inputJSON := `{"username": "keith", "email": "test@test.com"}`
        
        mockSvc.On("CreateUser", mock.Anything).Return(&User{ID: 1, UserFields: UserFields{Username: &username}}, nil).Once()

        req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(inputJSON))
        rr := httptest.NewRecorder()

        // Manually invoke the wrapped handler logic
        handler.Post(rr, req, User{UserFields: UserFields{Username: &username}})

        assert.Equal(t, http.StatusCreated, rr.Code)
    })
}