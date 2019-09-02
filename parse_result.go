package main

import (
	"fmt"
	"sync"
)

type ParseResult struct {
	mu         sync.Mutex
	count      int
	errorCount int
	details    map[string]int
}

func (t *ParseResult) AddError() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.errorCount++
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

	result := fmt.Sprintf("Total: %d\n", t.count)

	if t.errorCount > 0 {
		result += fmt.Sprintf("Finished with errors: %d urls", t.errorCount)
	}

	return result
}
