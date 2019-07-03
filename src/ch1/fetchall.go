// FetchAll fetches URLS in parallel and reports their times ad sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// go run fetchall.go https://github.com https://baidu.com
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)  // start a goroutine
		// goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。
		// main函数也是运行在一个goroutine中，而 go function 则表示创建一个新的goroutine，
		// 并在这个新的goroutine里执行这个函数。
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)  // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)  // send to channel ch
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()  // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nBytes, url)
}