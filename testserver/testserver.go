package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var requestCount uint64

func handler(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddUint64(&requestCount, 1)

	// Extract client IP details
	remoteAddr := r.RemoteAddr
	xff := r.Header.Get("X-Forwarded-For")
	if xff == "" {
		xff = "None"
	}

	// Print request info to terminal
	fmt.Printf("[%d] Request from Socket: %-21s | X-Forwarded-For: %s\n",
		count, remoteAddr, xff)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Printf("Test server listening on http://127.0.0.1%s\n\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
