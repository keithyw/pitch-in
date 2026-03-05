package auth

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/response"
)

type AuthHandler struct {
	svc AuthService
	jwt *jwt.JWTService
	log *slog.Logger
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Refresh string `json:"refresh"`
	User *users.User `json:"user"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh" validate:"required"`
}

type RefreshResponse struct {
	Token string `json:"token"`
	Refresh string `json:"refresh"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Email string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=255"`
	LastName string `json:"last_name" validate:"required,min=2,max=255"`
	Password string `json:"password" validate:"required,min=6,max=72,password_complex"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func NewAuthHandler(svc AuthService, jwt *jwt.JWTService, log *slog.Logger) *AuthHandler {
	return &AuthHandler{
		svc: svc,
		jwt: jwt,
		log: log,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request, req LoginRequest) {
	user, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := h.jwt.CreateJWT(user.ID, *user.Username)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Could not generate JWT: %s", err.Error()))
		return
	}

	refreshToken, err := h.jwt.CreateRefreshJWT(user.ID, *user.Username)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Could not generate JWT Refresh: %s", err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, LoginResponse{
		Refresh: refreshToken,
		Token: token,
		User: user,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request, req RefreshRequest) {
	claims, err := h.jwt.ParseJWT(req.Refresh)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	token, err := h.jwt.CreateJWT(claims.UserID, claims.Username)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	refresh, err := h.jwt.CreateRefreshJWT(claims.UserID, claims.Username)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	response.JSON(w, http.StatusOK, RefreshResponse{
		Token: token,
		Refresh: refresh,
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request, req RegisterRequest) {
	u := users.User{
		UserFields: users.UserFields{
			Username: &req.Username,
			Email: &req.Email,
			FirstName: &req.FirstName,
			LastName: &req.LastName,
		},
	}

	newUser, err := h.svc.Register(u, req.Password)	
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Registration failed: %s", err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, newUser)
}