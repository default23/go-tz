package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUrlContent(u string) string {
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
