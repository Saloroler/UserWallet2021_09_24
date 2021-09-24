package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func printHTTPResult(w http.ResponseWriter, resultHTTPCode int, jsonAnswer interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if jsonAnswer == nil && resultHTTPCode == http.StatusInternalServerError {
		body, _ := json.Marshal(map[string]string{"error": "Internal server error"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	body, err := json.Marshal(jsonAnswer)
	if err != nil {
		log.Print("error: failed to marshal json, error: ", err.Error())
		body, _ = json.Marshal(map[string]string{"error": "Internal server error"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	if resultHTTPCode >= 400 {
		body, _ = json.Marshal(map[string]interface{}{"error": jsonAnswer})
		w.WriteHeader(resultHTTPCode)
		w.Write(body)
		return
	}

	w.WriteHeader(resultHTTPCode)
	w.Write(body)
}
