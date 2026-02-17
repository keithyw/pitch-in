package users

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type UserHandler struct {
	svc UserService
	log *slog.Logger
}

func NewUserHandler(svc UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{
		svc: svc,
		log: log,
	}
}

func (h *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "userID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed get userID: %s", err.Error()))
		return
	}
	err = h.svc.DeleteUser(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete user: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *UserHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(User{}, h.log)
	filter, err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}
	users, err := h.svc.FindUserBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding users %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func (h *UserHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "userID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed get userID: %s", err.Error()))
		return
	}	
	user, err := h.svc.GetUser(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusNotFound, fmt.Sprintf("Failed retrieving user: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) Post(w http.ResponseWriter, req *http.Request, user User) {
	newUser, err := h.svc.CreateUser(user)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create user failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newUser)
}

func (h *UserHandler) Patch(w http.ResponseWriter, req *http.Request, userRequest PatchUserRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "userID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get userID: %s", err.Error()))
		return
	}
	updatedUser, err := h.svc.UpdateUser(*userRequest.ToModel(id))
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Update user failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, updatedUser)
}