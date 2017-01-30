package main

import (
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http.example.com")
	//...
	req, err := http.NewRequest("GET", "http://example.com", nil)
	//...
	req.Header.Add("User-Agent", "Our Custom User-Agent")
	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
	resp, err = client.Do(req)
}

func redirectPolicyFunc() {}
