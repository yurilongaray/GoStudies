package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// https://www.udemy.com/course/curso-go/learn/lecture/8758168#overview
func main() {
	t1 := requestTitle("https://www.google.com", "https://www.amazon.com")
	t2 := requestTitle("https://www.udemy.com", "https://www.youtube.com")

	fmt.Println("1°", <-t1, "- AND -", <-t2)
	fmt.Println("2°", <-t1, "- AND -", <-t2)

	champion := getTheFaster("https://www.google.com", "https://www.amazon.com", "https://www.udemy.com")

	fmt.Println(champion)
}

func getTheFaster(url1, url2, url3 string) string {

	c1 := requestTitle(url1)
	c2 := requestTitle(url2)
	c3 := requestTitle(url3)

	// select is a kind of "switch" but is only used for concurrency
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "They took too long"
	}
}

func requestTitle(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
