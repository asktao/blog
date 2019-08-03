package controllers

import (
	"encoding/json"
	"net/http"
)

func Message(status int, message string) (map[string]interface{}) {
	return map[string]interface{} {"status": status, "message": message}
}

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	JsonResponse(w, code, map[string]string{"message": message})
}
