package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// будем получать список адресов из stdin
	result := ParseUrls(os.Stdin, func(url, searchTerm string) (int, error) {
		// при вызове кол-бэка после получения адреса из stdin
		// делаем HTTP запрос и отдаем количество повторяющихся
		// searchTerm внутри этого контента
		content, err := GetUrlContent(url)

		if err != nil {
			return 0, err
		}

		return strings.Count(content, searchTerm), nil
	})

	fmt.Print(result)
}
