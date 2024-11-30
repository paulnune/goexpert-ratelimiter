package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("hello <-")
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
