package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,               // HLxxx
			IdleConnTimeout:    30 * time.Second, // HLxxx
			DisableCompression: true,             // HLxxx
		},
	}

	req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)

	resp, _ := c.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
}
