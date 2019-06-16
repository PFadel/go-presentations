package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * time.Duration(30), // HLxxx
	}

	req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req) // HLxxx
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
}
