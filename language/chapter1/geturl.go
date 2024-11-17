package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getUrl() {
	for _, url := range os.Args[1:] {
		//匹配前缀
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := io.ReadAll(resp.Body)
		written, err := io.Copy(os.Stdout, resp.Body) //避免申请了一个缓冲区b来存储整个响应数据
		fmt.Printf("%d\n", written)
		fmt.Println(resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
