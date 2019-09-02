package utils_test

import (
	"github.com/default23/go_tz/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetUrls(t *testing.T) {

	result := []string{"https://google.com", "https://fb.com"}

	tests := map[string]struct {
		content string
	}{
		"should parse urls from the reader": {
			content: "https://google.com\nhttps://fb.com",
		},
		"should ignore invalid urls": {
			content: "https://google.com\nhttps://fb.com\nfb.ru\none\n123\nhello\nhttp//hello.world",
		},
		"should ignore next urls if received stop word": {
			content: "https://google.com\nhttps://fb.com\nstop_word\nhttps://fb.ru\none\n123\nhello\nhttp//hello.world",
		},
	}

	for name, _t := range tests {
		t.Log(name)
		reader := strings.NewReader(_t.content)

		assert.Equal(t, result, utils.GetUrls(reader))
	}
}
