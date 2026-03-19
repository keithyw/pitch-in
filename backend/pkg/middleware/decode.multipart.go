package middleware

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/keithyw/pitch-in/pkg/response"
)

var decoder = schema.NewDecoder()

func DecodeMultipart[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			response.ErrorJSON(w, http.StatusBadRequest, "Form submission to large")
			return
		}

		var payload T

		if err := decoder.Decode(&payload, r.PostForm); err != nil {
			response.ErrorJSON(w, http.StatusBadRequest, "Invalid form submission")
			return
		}

		if err := validate.Struct(payload); err != nil {
			response.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		next(w, r, payload)
	}
}