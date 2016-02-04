// Fetch prints the content found at a URL.

package main

import (
	"fmt"
//	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
//		fmt.Printf("url: %s\n", url)
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
//		fmt.Printf("url1: %s\n", url)

		resp, err := http.Get(url)
		fmt.Printf("Status: %s\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

//		b, err := ioutil.ReadAll(resp.Body)
		numBytes, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("num bytes read: %d\n", numBytes)
//		fmt.Printf("%s", b)
	}
}
