package media

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/pkg/response"
)

type MediaHandler struct {
	svc MediaService
	log *slog.Logger
}

func NewMediaHandler(svc MediaService,log *slog.Logger) *MediaHandler {
	return &MediaHandler{
		svc: svc,
		log: log,
	}
}

func (h *MediaHandler) Get(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	data := make(map[string]any)
	for k, v := range params {
		data[k] = v[0]
	}
	res, err := h.svc.QueryShow(req.Context(), data)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, res)
}