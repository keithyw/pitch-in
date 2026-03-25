package tags

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type AttachEntityRequest struct {
	EntityID int64 `json:"entity_id" validate:"required"`
	Entity string `json:"entity" validate:"required"`
}

type DetachEntityRequest struct {
	EntityID int64 `json:"entity_id" validate:"required"`
	Entity string `json:"entity" validate:"required"`
}

type TagHandler struct {
	svc TagService
	log *slog.Logger
}

func NewTagHandler(svc TagService, log *slog.Logger) *TagHandler {
	return &TagHandler{
		svc: svc,
		log: log,
	}
}

func (h *TagHandler) AttachEntity(w http.ResponseWriter, req *http.Request, entityRequest AttachEntityRequest) {
	tagID, err := strconv.ParseInt(chi.URLParam(req, "tagID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse tagID: %s", err.Error()))
		return
	}
	err = h.svc.AttachEntity(entityRequest.EntityID, tagID, entityRequest.Entity)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to attach entity: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, nil)
}

func (h *TagHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "tagID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get tagID: %s", err.Error))
		return
	}
	err = h.svc.DeleteTag(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete tag: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *TagHandler) DetachEntity(w http.ResponseWriter, req *http.Request, entityRequest DetachEntityRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "tagID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse tagID: %s", err.Error()))
		return
	}
	err = h.svc.DetachEntity(entityRequest.EntityID, id, entityRequest.Entity)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to detach entity: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}


func (h *TagHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(Tag{}, h.log)
	filter, err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}
	tags, err := h.svc.FindTagsBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding tags: %s", err.Error()))
		return
	}

	count, err := h.svc.CountTags(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed counting tags: %s", err.Error()))
		return
	}
	response.PaginatedJSON(w, http.StatusOK, count, tags)
}

func (h *TagHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "tagID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get tagID: %s", err.Error()))
		return
	}
	tag, err := h.svc.GetTag(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed retrieving tag: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, tag)
}

func (h *TagHandler) Post(w http.ResponseWriter, req *http.Request, tag Tag) {
	newTag, err := h.svc.CreateTag(tag)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create tag failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newTag)
}

func (h *TagHandler) Patch(w http.ResponseWriter, req *http.Request, tagRequest PatchTagRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "tagID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to get tagID: %s", err.Error()))
		return
	}
	updatedTag, err := h.svc.UpdateTag(*tagRequest.ToModel(id))
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Update tag failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, updatedTag)
}