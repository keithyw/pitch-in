package response

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice && val.IsNil() {
		data = reflect.MakeSlice(val.Type(), 0, 0).Interface()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ErrorJSON(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}