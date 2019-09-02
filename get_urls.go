package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"sync"
)

// Ищет количество вхождений searchTerm в контент
// полученный по url
type ContentParser func(url, searchTerm string) (int, error)

// Получает из reader'а список адресов и формирует результат
// вычисления количества повторяющихся "Go"
func ParseUrls(reader io.Reader, count ContentParser) ParseResult {
	scanner := bufio.NewScanner(reader)
	result := ParseResult{}

	// Ограничиваем количество максимальных потоков
	// с помощью паттерна семафор до 5
	s := make(chan int, 5)
	var wg sync.WaitGroup

	for scanner.Scan() {
		// Будет блокировать дальнейшее выполнение, если канал заполнен (k=5)
		s <- 1
		text := scanner.Text()
		wg.Add(1)

		// Останавливаем чтение из reader'a
		// сделал для себя, что бы удобнее было отлаживать через debugger
		// на условие задания не влияет
		if text == "stop_word" {
			wg.Done()
			break
		}

		// необходимо проверить, что полученная строка
		// является валидным URL адресом
		if _, err := url.ParseRequestURI(text); err != nil {
			fmt.Printf("provided url is not valid: %s, will not be parsed (should be like 'http://some.url')\n", text)
			continue
		}

		go func(url string) {
			defer wg.Done()

			if occurrences, err := count(url, "Go"); err != nil {
				result.AddError()
			} else {
				result.Add(text, occurrences)
				fmt.Printf("Count for %s: %d\n", url, occurrences)
			}

			<-s
		}(text)
	}

	wg.Wait()
	return result
}
