package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// будем получать список адресов из stdin
	result := ParseUrls(os.Stdin, func(url, searchTerm string) int {
		// при вызове кол-бэка после получения адреса из stdin
		// делаем HTTP запрос и отдаем количество повторяющихся
		// searchTerm внутри этого контента
		content := GetUrlContent(url)
		return strings.Count(content, searchTerm)
	})

	fmt.Print(result)
}
