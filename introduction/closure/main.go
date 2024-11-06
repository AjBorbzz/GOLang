package main

import (
	"fmt"
	"sync"
	"time"
)

func NewRateLimiter(requestsPerSecond int) func() bool {
	var lastTime time.Time
	var mu sync.Mutex

	return func() bool {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if lastTime.IsZero() || now.Sub(lastTime) >= time.Second/time.Duration(requestsPerSecond) {
			lastTime = now
			return true
		}
		return false
	}
}

func main() {
	// Create a rate limiter that allows 5 requests per second
	limiter := NewRateLimiter(5)

	// Simulate making API calls
	for i := 0; i < 10; i++ {
		if limiter() {
			fmt.Printf("Request %d allowed\n", i+1)
		} else {
			fmt.Printf("Request %d denied\n", i+1)
		}
		time.Sleep(100 * time.Millisecond) // Simulate a slight delay between requests
	}
}
