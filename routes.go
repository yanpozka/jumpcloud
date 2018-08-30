package main

import "net/http"

// makeRoutes abstracts the process of registering routes returning a plain http.Handler,
// switching of router or middleware chains is easier in this way.
func makeRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/hash", loggerPanic(
		saveStatsAndLogger(httpMethod(http.MethodPost, http.HandlerFunc(handleHash)))))

	mux.Handle("/shutdown", loggerPanic(httpMethod(http.MethodGet, http.HandlerFunc(handleShutdown))))

	mux.Handle("/stats", loggerPanic(httpMethod(http.MethodGet, http.HandlerFunc(handleStats))))

	return mux
}
