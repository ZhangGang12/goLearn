package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetchall() {
	start := time.Now()
	ch := make(chan string) //建立一个双向通道
	for i, url := range os.Args[1:] {
		go fetchV1(url, ch, fmt.Sprintf("getUrl%d", i)) // start a goroutine
		time.Sleep(500 * time.Millisecond)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) { //行参将双向通道隐式转化成只接受的单项通道
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // nbytes 返回的是接受的字节数
	resp.Body.Close()                                 // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 统计花费的时间
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func fetchV1(url string, ch chan<- string, filename string) { //行参将双向通道隐式转化成只接受的单项通道
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	//创建一个空文件，用来存储返回的数据
	f, err := os.Create("./" + filename)
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("create file failed: %s", err)
		return
	}
	defer resp.Body.Close()

	// 将响应头写入文件
	headerWriter := bufio.NewWriter(f)
	for key, values := range resp.Header {
		for _, value := range values {
			headerWriter.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
	}
	headerWriter.Flush()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 统计花费的时间
	ch <- fmt.Sprintf("%.2fs  %7d %s %s", secs, nbytes, resp.Status, url)
}
