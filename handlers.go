package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/yanpozka/jumpcloud/hashing"
)

type statsResponse struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

// TODO: create a Server struct and remove this global variables and move them inside Server
var (
	stats   statsResponse
	statsMX = new(sync.RWMutex)
)

func handleHash(w http.ResponseWriter, r *http.Request) {
	passwd := r.PostFormValue("password")
	if passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	time.Sleep(5 * time.Second)

	hash := hashing.HashBase64(passwd)
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(hash)); err != nil {
		log.Printf("Error writing response: %q", err)
	}
}

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	shutdownCH <- struct{}{}

	w.WriteHeader(http.StatusOK)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	statsMX.RLock()
	res := statsResponse{
		Total:   stats.Total,
		Average: stats.Average,
	}
	statsMX.RUnlock()

	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when encoding response: %q", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
