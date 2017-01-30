package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	//...
	req.Header.Add("User-Agent", "GoBook Custom User-Agent")
	//...
	client := &http.Client{}
	resp, err := client.Do(req)
	//...
}
