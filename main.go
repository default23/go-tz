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
