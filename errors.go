package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorMessage struct {
	Error_msg string
}

func FileNumberNotPresentError(w http.ResponseWriter, r *http.Request, err error) errorMessage {
	res := Error(w, r, http.StatusBadRequest, err)

	return res
}

func Error(w http.ResponseWriter, r *http.Request, statusCode int, err error) errorMessage {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-type", "application/json")
	res := errorMessage{
		Error_msg: err.Error(),
	}
	return res
}

func ReturnJSONErrResponse(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	em := errorMessage{
		Error_msg: err.Error(),
	}
	w.WriteHeader(statusCode)
	responseJSON, err := json.Marshal(em)
	if err != nil {
		log.Println("Error decoding JSON")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}
