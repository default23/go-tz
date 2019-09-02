package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
)

// Получает из reader'а список адресов
func GetUrls(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	var urls []string
	for scanner.Scan() {
		raw := scanner.Text()

		if raw == "stop_word" {
			break
		}
		// необходимо проверить, что полученная строка
		// является валидным URL адресом
		if _, err := url.ParseRequestURI(raw); err != nil {
			fmt.Printf("provided url is not valid: %s, (should be like 'http://some.url'), will not be used in the next operations \n", raw)
			continue
		}

		urls = append(urls, scanner.Text())
	}

	return urls
}
