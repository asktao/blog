package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Status		int				`json:"status"`
	Message		string			`json:"message"`
	Data		interface{}		`json:"data"`
}

type Controller struct {

}

type ControllerInterface interface {
	JsonResponse(w http.ResponseWriter, response Response)
	Pagination (r *http.Request) (limit uint64, offset uint64)
}

func (c *Controller) Pagination (r *http.Request) (limit uint64, offset uint64) {
	limit, _ = strconv.ParseUint(r.FormValue("limit"), 10, 64)

	offset, _ = strconv.ParseUint(r.FormValue("offset"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	return limit, offset
}

func (c *Controller) JsonResponse(w http.ResponseWriter, response Response) {
	reps, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF=8")
	w.WriteHeader(response.Status)
	w.Write(reps)
	return
}