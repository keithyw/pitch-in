package users

import (
	"fmt"
	"log/slog"
	"net/http"

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

func (h *UserHandler) Post(w http.ResponseWriter, req *http.Request, user User) {
	newUser, err := h.svc.CreateUser(user)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create User Failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newUser)
	
}