package main

import (
	"fmt"
	"sync"
)

type total struct {
	mu      sync.Mutex
	count   int
	details map[string]int
}

func (t *total) Add(url string, count int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.count += count
	t.details[url] = count
}

func (t total) String() string {
	var details string
	for k, v := range t.details {
		details += fmt.Sprintf("Count for %s: %d\n", k, v)
	}

	return fmt.Sprintf("%sTotal: %d\n", details, t.count)
}
