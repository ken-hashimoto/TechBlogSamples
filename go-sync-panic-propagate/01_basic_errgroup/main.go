package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://example.com",
	}

	for _, url := range urls {
		url := url // ループ変数のキャプチャ
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	} else {
		fmt.Println("All requests succeeded!")
	}
}
