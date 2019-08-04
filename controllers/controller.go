package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status		uint16			`json:"status"`
	Message		string			`json:"message"`
	Data		interface{}		`json:"data"`
}

type Controller struct {

}

type ControllerInterface interface {
	JsonResponse(w http.ResponseWriter, code int, payload interface{})
}

func (c Controller) JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF=8")
	w.WriteHeader(code)
	w.Write(response)
	return
}

//func ErrorResponseWithLog(w http.ResponseWriter, code int, message string, err error) {

//	JsonResponse(w, code, map[string]string{"message": message})
//	return
//}
