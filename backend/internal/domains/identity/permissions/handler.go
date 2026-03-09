package permissions

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type PermissionHandler struct {
	svc PermissionService
	log *slog.Logger
}

func NewPermissionHandler(svc PermissionService, log *slog.Logger) *PermissionHandler {
	return &PermissionHandler{
		svc: svc,
		log: log,
	}
}

func (h *PermissionHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "permissionID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to getpermissionID: %s", err.Error()))
		return
	}
	err = h.svc.DeletePermission(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete permission: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *PermissionHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(Permission{}, h.log)
	filter, err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}
	permissions, err := h.svc.FindPermissionBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding permissions: %s", err.Error()))
		return
	}

	count, err := h.svc.CountPermissions(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Faile permission count: %s", err.Error()))
		return
	}
	response.PaginatedJSON(w, http.StatusOK, count, permissions)
}

func (h *PermissionHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "permissionID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get permissionID: %s", err.Error()))
		return
	}
	perm, err := h.svc.GetPermission(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed retrieving permission: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, perm)
}

func (h *PermissionHandler) Post(w http.ResponseWriter, req *http.Request, permission Permission) {
	newPerm, err := h.svc.CreatePermission(permission)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create permission failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newPerm)
}

func (h *PermissionHandler) Patch(w http.ResponseWriter, req *http.Request, permRequest PatchPermissionRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "permissionID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get permissionID: %s", err.Error()))
		return
	}
	updatedPerm, err := h.svc.UpdatePermission(*permRequest.ToModel(id))
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Update permission failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, updatedPerm)
}