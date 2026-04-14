package identifier

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/pkg/response"
)

type IdentifyImageRequest struct {
	Prompt string `schema:"prompt" validate:"required,max=500"`
}

type IdentifierHandler struct {
	svc IdentifierService
	log *slog.Logger
}

func NewIdentifierHandler(svc IdentifierService, log *slog.Logger) *IdentifierHandler {
	return &IdentifierHandler{
		svc: svc,
		log: log,
	}
}

func (h *IdentifierHandler) Post(w http.ResponseWriter, req *http.Request, iReq IdentifyImageRequest) {
	file, header, err := req.FormFile("file")
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, "File is required")
		return
	}
	defer file.Close()

	res, err := h.svc.IdentifyImage(req.Context(), iReq.Prompt, header.Filename, file)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, res)
}