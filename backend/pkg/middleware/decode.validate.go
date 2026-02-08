package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/keithyw/pitch-in/pkg/response"
)

var validate = validator.New()

func DecodeAndValidate[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var payload T
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			response.ErrorJSON(w, http.StatusBadRequest, "Malformed request")
			return
		}
		defer r.Body.Close()

		if err := validate.Struct(payload); err != nil {
			response.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())	
			return
		}
		next(w, r, payload)
	}
}