package main

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

const startTimeKey = "__startTimeKey"

func httpMethod(method string, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//
		next.ServeHTTP(w, r)
	})
}

func loggerPanic(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			val := recover()
			if val == nil {
				return
			}
			log.Printf("Recovering from panic with value: '%v'", val)
			debug.PrintStack()
		}()

		start := time.Now()
		ctx := context.WithValue(r.Context(), startTimeKey, start)

		//
		next.ServeHTTP(w, r.WithContext(ctx))

		delta := time.Since(start)
		logRequest(r, delta)
	})
}

func saveStatsAndLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		next.ServeHTTP(w, r)

		start, isTime := r.Context().Value(startTimeKey).(time.Time)
		if !isTime {
			log.Panic("Start time not found :(")
			return
		}

		statsMX.Lock()
		// 10e3 : from nano 10e9 to micro 10e6
		deltaMicro := int64(time.Since(start)) / 1e3

		if stats.Total == 0 {
			stats.Average = deltaMicro
		} else {
			currentAvg := stats.Average
			// incremental average
			currentAvg += (deltaMicro - currentAvg) / stats.Total
			stats.Average = currentAvg
		}

		stats.Total++
		// log.Printf("Saved stats: %+v start: %v", stats, start)
		statsMX.Unlock()
	})
}

func logRequest(r *http.Request, d time.Duration) {
	log.Printf("[%s]: %s | Time consumed: %v", r.Method, r.RequestURI, d)
}
