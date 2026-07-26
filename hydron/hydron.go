package hydron

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func generateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(223)+1,
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(254)+1)
}

func Create_packetjob(
	src string,
	dest string,
	prt int,
	path string,
	wrkrs int,
	reqs int,
	dur time.Duration,
	fkIPs bool) *PacketJob {
	return &PacketJob{
		src:         src,
		dst:         dest,
		port:        prt,
		path:        path,
		workers:     wrkrs,
		requests:    reqs,
		duration:    dur,
		SimulateIPs: fkIPs,
	}
}

func Hydron_Creator(job PacketJob) *Hydron {

	sock_man := &http.Transport{
		MaxIdleConnsPerHost: job.workers,
		MaxConnsPerHost:     job.workers,
		IdleConnTimeout:     30 * time.Second,
	}

	return &Hydron{
		config: job,
		client: &http.Client{
			Timeout:   5 * time.Second,
			Transport: sock_man,
		},
	}
}

func (hydron *Hydron) run_request(
	ctx context.Context,
	wait_group *sync.WaitGroup,
	jobs chan struct{},
	results chan struct{ ok bool }) {

	defer wait_group.Done()

	target := fmt.Sprintf("http://%s:%d%s", hydron.config.dst, hydron.config.port, hydron.config.path)

	for range jobs {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)

		if err != nil {
			results <- struct{ ok bool }{ok: false}
			continue
		}

		if hydron.config.SimulateIPs {
			fakeIP := generateRandomIP()
			req.Header.Set("X-Forwarded-For", fakeIP)
			req.Header.Set("X-Real-IP", fakeIP)
		}

		response, err := hydron.client.Do(req)
		if err != nil {
			results <- struct{ ok bool }{ok: false}
			continue
		}

		results <- struct{ ok bool }{ok: response.StatusCode < 400}
		response.Body.Close()

	}
}

func (hydron *Hydron) Run(ctx context.Context) {
	jobs := make(chan struct{}, hydron.config.requests)
	results := make(chan struct{ ok bool }, hydron.config.requests)

	var wait_group sync.WaitGroup

	for i := 0; i < hydron.config.workers; i++ {
		wait_group.Add(1)
		go hydron.run_request(ctx, &wait_group, jobs, results)
	}

	go func() {
		for res := range results {
			atomic.AddUint64(&hydron.stats.totalRequests, 1)

			if res.ok {
				atomic.AddUint64(&hydron.stats.successCount, 1)
			} else {
				atomic.AddUint64(&hydron.stats.errorCount, 1)
			}
		}
	}()

	ticker := time.NewTicker(time.Second / time.Duration(hydron.config.requests))
	defer ticker.Stop()

	stopTimer := time.NewTimer(hydron.config.duration)
	defer stopTimer.Stop()

Loop:
	for {
		select {
		case <-stopTimer.C:
			break Loop
		case <-ctx.Done():
			break Loop
		case <-ticker.C:
			select {
			case jobs <- struct{}{}:
			default:
			}
		}
	}

	close(jobs)
	wait_group.Wait()
	close(results)

	fmt.Printf("Done. Total Requests: %d, Success: %d, Errors: %d\n",
		atomic.LoadUint64(&hydron.stats.totalRequests),
		atomic.LoadUint64(&hydron.stats.successCount),
		atomic.LoadUint64(&hydron.stats.errorCount))

}
