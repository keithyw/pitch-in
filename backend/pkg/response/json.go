package response

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type ResponseCollection struct {
	Count int64 `json:"count"`
	Results interface{} `json:"results"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice && val.IsNil() {
		data = reflect.MakeSlice(val.Type(), 0, 0).Interface()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func PaginatedJSON(w http.ResponseWriter, status int, count int64, data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice && val.IsNil() {
		data = reflect.MakeSlice(val.Type(), 0, 0).Interface()
	}

	payload := ResponseCollection{
		Count: count,
		Results: data,
	}

	JSON(w, status, payload)
}

func ErrorJSON(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}