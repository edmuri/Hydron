package hydron

import (
	"net/http"
	"time"
)

type PacketJob struct {
	src         string
	dst         string
	port        int
	workers     int
	requests    int
	duration    time.Duration
	SimulateIPs bool
}

type Hydron struct {
	config PacketJob
	client *http.Client
	stats  struct {
		totalRequests uint64
		successCount  uint64
		errorCount    uint64
	}
}
