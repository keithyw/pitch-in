package roles

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type RoleHandler struct {
	svc RoleService
	log *slog.Logger
}

func NewRoleHandler(svc RoleService, log *slog.Logger) *RoleHandler {
	return &RoleHandler{
		svc: svc,
		log: log,
	}
}

func (h *RoleHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "roleID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get roleID: %s", err.Error()))
		return
	}
	err = h.svc.DeleteRole(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete role: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *RoleHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(Role{}, h.log)
	filter, err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}

	roles, err := h.svc.FindRoleBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding roles: %s", err.Error()))
		return
	}

	count, err := h.svc.CountRoles(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed role count: %s", err.Error()))
		return
	}
	response.PaginatedJSON(w, http.StatusOK, count, roles)
}

func (h *RoleHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "roleID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse roleID: %s", err.Error()))
		return
	}
	role, err := h.svc.GetRole(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get role: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, role)
}

func (h *RoleHandler) Post(w http.ResponseWriter, req *http.Request, role Role) {
	newRole, err := h.svc.CreateRole(role)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create role failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newRole)
}

func (h *RoleHandler) Patch(w http.ResponseWriter, req *http.Request, roleRequest PatchRoleRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "roleID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse roleID: %s", err.Error()))
		return
	}
	updatedRole, err := h.svc.UpdateRole(*roleRequest.ToModel(id))
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Update role failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, updatedRole)
}