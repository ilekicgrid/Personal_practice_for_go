package main

import (
	"io/ioutil"
	"log"
	"net/http"
    "fmt"
)

func main() {
	client := *http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create request")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Unable to get response")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}

    fmt.Println(string (content))
}
