package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	data := struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{
		Status: "ok",
		Data:   "server running",
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonBytes, _ := json.Marshal(data)
		_, _ = w.Write(jsonBytes)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
