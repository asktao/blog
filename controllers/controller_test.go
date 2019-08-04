package controllers

import (
	"testing"
)

type Columns struct {
	One string `json:"one"`
	Two string `json:"two"`
}

func TestMessage(t *testing.T) {

}

//func TestJsonResponse(t *testing.T) {
//	var err error
//	h := http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
//		JsonResponse(w, 200, Columns{"hello", "world"})
//	})
//
//	res := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/n", nil)
//	h.ServeHTTP(res, req)
//
//	expectNil(t, err)
//
//	expect(t, res.Code, 200)
//	expect(t, res.Header().Get(ContentType), ContentJSON+"; charset=UTF-8")
//	expect(t, res.Body.String(), "{\"one\":\"hello\",\"two\":\"world\"}")
//}


func TestErrorResponse(t *testing.T) {

}