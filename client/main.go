package main

import (
	"fmt"
	"io"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "localhost:8080"
	uStr := u.String()

	rsp, err := retryablehttp.Get(uStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()

	body, _ := io.ReadAll(rsp.Body)
	fmt.Println(string(body))
}
