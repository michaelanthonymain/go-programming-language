package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		u := appendHTTP(url)
		resp, err := http.Get(u)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:  %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("Status: %s\n", resp.Status)
	}
}

func appendHTTP(u string) string {
	if strings.HasPrefix(u, "http://") {
		return u
	}

	return "http://" + u
}
