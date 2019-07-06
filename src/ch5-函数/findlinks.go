package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

// It can not work, sorry.

// 一个函数可以返回多个值。许多标准库中的函数返回2个值，一个是期望值，一个是函数出错是的错误信息

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			continue
		}

		for _, link := range links {
			fmt.Printf(link)
		}
	}
}

func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)

	}

	return visit(nil, doc), nil
}
