package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	retryClient := retryablehttp.NewClient()
	standardCliend := retryClient.StandardClient()
	retryClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp.StatusCode == http.StatusServiceUnavailable {
			return false, nil
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	req, err := http.NewRequest(
		"GET",
		"http://localhost:8080",
		nil,
	)

	rsp, err := standardCliend.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()

	body, _ := io.ReadAll(rsp.Body)
	fmt.Println(string(body))
}
