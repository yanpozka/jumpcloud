package main

import (
	"net/http"
	"time"

	"github.com/yanpozka/jumpcloud/hashing"
)

func handleHash(w http.ResponseWriter, r *http.Request) {
	passwd := r.PostFormValue("password")
	if passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	time.Sleep(5 * time.Second)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hashing.HashBase64(passwd)))
}
