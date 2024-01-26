package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

func FileLineSearch(w http.ResponseWriter, r *http.Request) {
	queryparams := r.URL.Query()

	n := queryparams.Get("n")
	m := queryparams.Get("m")

	if n == "" {
		log.Println("The file number has not been provided")

		res := FileNumberNotPresentError(w, r, errors.New("file Number Not present"))

		responseJSON, err := json.Marshal(res)

		if err != nil {
			log.Println("Error decoding JSON")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		w.Write(responseJSON)
		return
	}

	filePath := "tmp/data/" + n + ".txt"
	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Println("No such File with the given name exists")
		ReturnJSONErrResponse(w, r, errors.New("no such File with the given name exists"), http.StatusBadRequest)
		return
	}

	file_main := string(content)

	if m == "" {
		//server entire file
		// w.Header().Set("Content-type", "plain/text")
		w.Write([]byte(file_main))
		return
	} else if m != "" && n != "" {
		//server line m of the file n

		if !NumberValidator(m) {
			log.Println("value of m is corrupted,error converting from string to int")
			ReturnJSONErrResponse(w, r, errors.New("query parameter m is not in accordance with the condition"), http.StatusBadRequest)
			return
		}

		length := cache[n+"_len"]
		m_int := ConvertDataToInt(m)
		if m_int > ConvertDataToInt(length) {
			log.Println("Value of m greater than the lines present in file")
			ReturnJSONErrResponse(w, r, errors.New("m greater than number of lines present in file"), http.StatusBadRequest)
			return
		}

		// w.Header().Set("Content-type", "plain/text")
		w.Write([]byte(cache[n+"_"+m]))
		return
	}
}

func SendResponse(w http.ResponseWriter, r *http.Request, data string) {
	// w.Header().Set("Content-type", "plain/text")
	w.Write([]byte(data))
}
