package main

import (
	"fmt"
	"github.com/default23/go_tz/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

/**
Программа читает из stdin строки, содержащие URL.
На каждый URL нужно отправить HTTP-запрос методом GET
и посчитать кол-во вхождений строки "Go" в теле ответа.
В конце работы приложение выводит на экран общее количество
найденных строк "Go" во всех переданных URL, например:

$ echo -e 'https://golang.org\nhttps://golang.org' | go run 1.go
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 18

Каждый URL должен начать обрабатываться сразу после вычитывания
и параллельно с вычитыванием следующего.
URL должны обрабатываться параллельно, но не более k=5 одновременно.
Обработчики URL не должны порождать лишних горутин, т.е. если k=5,
а обрабатываемых URL-ов всего 2, не должно создаваться 5 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартную библиотеку.
*/

func main() {
	urls := utils.GetUrls(os.Stdin)
	var result total
	var wg sync.WaitGroup

	result.details = map[string]int{}
	wg.Add(len(urls))

	for _, u := range urls {

		go (func(url string) {
			defer wg.Done()
			content := getUrlContent(url)
			result.Add(url, strings.Count(content, "Go"))
		})(u)
	}

	wg.Wait()
	fmt.Print(result)
}

func getUrlContent(u string) string {
	res, err := http.Get(u)
	if err != nil {
		fmt.Printf("Error while performing GET request: %s", u)
		return ""
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read response body for url: %s; error: %v", u, err)
		return ""
	}

	return string(content)
}
