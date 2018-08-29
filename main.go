package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	port := getEnvStr("PORT", "8080")

	srv := &http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`hello amigoooosss`))
		}),

		ReadTimeout:       getEnvDuration("READ_TIMEOUT_SECONDS", 10) * time.Second,
		ReadHeaderTimeout: getEnvDuration("READHEADER_TIMEOUT_SECONDS", 5) * time.Second,
		WriteTimeout:      getEnvDuration("WRITE_SECONDS", 15) * time.Second,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)

	go func() {
		log.Printf("Serving on port :%s ...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	timeout := getEnvDuration("SHUTDOWN_TIMEOUT_SECONDS", 11) * time.Second

	// blocks until we get a terminal OS signal
	//
	osSignal := <-ch

	log.Printf("Got OS signal: '%v', shuting down the server with timeout: %v ", osSignal, timeout)

	srv.SetKeepAlivesEnabled(false)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown the server: %v", err)
	}
}

func getEnvStr(name, defaultVal string) string {
	envVal := os.Getenv(name)
	if envVal == "" {
		return defaultVal
	}
	return envVal
}

func getEnvDuration(name string, defaultVal int) time.Duration {
	envVal := os.Getenv(name)
	if envVal == "" {
		return time.Duration(defaultVal)
	}
	num, _ := strconv.Atoi(envVal)
	return time.Duration(num)
}
