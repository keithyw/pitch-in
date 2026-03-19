package assets

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/response"
)

type CreateAssetRequest struct {
	ObjectKey string `schema:"object_key" validate:"required,max=500"`
	Width *int `schema:"width"`
	Height *int `schema:"height"`
}

type AssetHandler struct {
	svc AssetService
	log *slog.Logger
}

func NewAssetHandler(svc AssetService, log *slog.Logger) *AssetHandler {
	return &AssetHandler{
		svc: svc,
		log: log,
	}
}

func (h *AssetHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "assetID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse assetID: %s", err.Error()))
		return
	}

	err = h.svc.DeleteAsset(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed deleting asset: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *AssetHandler) FindBy(w http.ResponseWriter, req *http.Request) {
	p := repository.NewParser(Asset{}, h.log)
	filter , err := p.Parse(req.URL.Query())
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}

	assets, err := h.svc.FindAssetBy(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed finding assets: %s", err.Error()))
		return
	}

	count, err := h.svc.CountAssets(*filter)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to count permissions: %s", err.Error()))
		return
	}

	response.PaginatedJSON(w, http.StatusOK, count, assets)
}

func (h *AssetHandler) Get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(req, "assetID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed parsing query: %s", err.Error()))
		return
	}
	asset, err := h.svc.GetAsset(id)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed getting asset: %s", err.Error()))
		return
	}
	response.JSON(w, http.StatusOK, asset)
}

func (h *AssetHandler) Post(w http.ResponseWriter, req *http.Request, assetReq CreateAssetRequest) {
	file, _, err := req.FormFile("file")
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, "File is required")
		return
	}
	defer file.Close()

	asset := Asset{
		AssetFields: AssetFields{
			ObjectKey: &assetReq.ObjectKey,
			Width:     assetReq.Width,
			Height:    assetReq.Height,
		},
	}

	newAsset, err := h.svc.CreateAsset(asset, file)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, newAsset)
}

func (h *AssetHandler) Patch(w http.ResponseWriter, req *http.Request, assetReq CreateAssetRequest) {
	id, err := strconv.ParseInt(chi.URLParam(req, "assetID"), 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, "Invalid asset ID")
		return
	}

	// Check if a new file was included in the multipart form
	var file io.ReadCloser
	f, _, err := req.FormFile("file")
	if err == nil {
		file = f
		defer file.Close()
	}

	// Map request to the Asset model
	asset := Asset{
		BaseModel: model.BaseModel{ID: id},
		AssetFields: AssetFields{
			ObjectKey: &assetReq.ObjectKey,
			Width:     assetReq.Width,
			Height:    assetReq.Height,
		},
	}

	updated, err := h.svc.UpdateAsset(asset, file)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, updated)
}