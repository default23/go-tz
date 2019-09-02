package main

import (
	"fmt"
	"sync"
)

type ParseResult struct {
	mu      sync.Mutex
	count   int
	details map[string]int
}

func (t *ParseResult) Add(url string, count int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.details == nil {
		t.details = map[string]int{}
	}

	t.count += count
	t.details[url] = count
}

func (t ParseResult) String() string {
	return fmt.Sprintf("Total: %d\n", t.count)
}
