package items

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type ItemHandler struct {
	svc ItemService
	log *slog.Logger
}

func NewItemHandler(svc ItemService, log *slog.Logger) *ItemHandler {
	return &ItemHandler{
		svc: svc,
		log: log,
	}
}

func (h *ItemHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "itemID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse itemID: %s", err.Error))
		return
	}
	err = h.svc.DeleteItem(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete item: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *ItemHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(Item{}, h.log)
	filter, err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}
	items, err := h.svc.FindItemsBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding items: %s", err.Error()))
		return
	}
	count, err := h.svc.CountItems(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed counting items: %s", err.Error()))
		return
	}
	response.PaginatedJSON(w, http.StatusOK, count, items)
}

func (h *ItemHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "itemID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse itemID: %s", err.Error()))
		return
	}
	item, err := h.svc.GetItem(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed retrieving item: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, item)
}

func (h *ItemHandler) Post(w http.ResponseWriter, req *http.Request, item Item) {
	newItem, err := h.svc.CreateItem(item)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Create item failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusCreated, newItem)
}

func (h *ItemHandler) Patch(w http.ResponseWriter, req *http.Request, itemRequest PatchItemRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "itemID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing itemID: %s", err.Error()))
		return
	}
	updatedItem, err := h.svc.UpdateItem(*itemRequest.ToModel(id))
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Update item failed: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, updatedItem)
}
	