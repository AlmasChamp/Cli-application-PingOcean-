package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	args := os.Args[1:]

	urls := strings.Split(args[1], ",")
	search := args[3]
	mp := make(map[string]int)

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	mutex := new(sync.Mutex)
	wg := &sync.WaitGroup{}
	ch := make(chan string, 3)

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go doIt(ctx, ch, wg, mutex, mp, search)
	}

	for _, url := range urls {
		ch <- url
	}

	close(ch)
	wg.Wait()
	fmt.Println("End")

}

func doIt(ctx context.Context, ch chan string, wg *sync.WaitGroup, mutex *sync.Mutex, mp map[string]int, search string) {

	url := <-ch
	// ИЛИ mutex.Lock()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(nil))
	if err != nil {
		log.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	mutex.Lock()
	mp[url] = strings.Count(string(body), search)
	mutex.Unlock()
	fmt.Println("<", url, ">", "-", "<", strings.Count(string(body), search), ">")

	wg.Done()
}

// FAN-IN

// func main() {

// 	args := os.Args[1:]

// 	urls := strings.Split(args[1], ",")
// 	search := args[3]
// 	mp := make(map[string]int)

// 	input := gen(urls)
// 	ch1 := sq(input)
// 	ch2 := sq(input)
// 	ch3 := sq(input)

// 	for n := range merge(search, mp, ch1, ch2, ch3) {
// 		fmt.Println(n)
// 	}
// 	fmt.Println("End")

// }

// func gen(urls []string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		for _, n := range urls {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(input <-chan string) <-chan string {

// 	out := make(chan string)
// 	go func() {
// 		for url := range input {
// 			out <- url
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func merge(search string, mp map[string]int, chs ...<-chan string) <-chan string {
// 	var wg sync.WaitGroup
// 	mutex := new(sync.Mutex)
// 	ctx := context.Background()
// 	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
// 	out := make(chan string)

// 	output := func(ch <-chan string) {
// 		for url := range ch {
// 			out <- doIt(ctx, mutex, url, search, mp)
// 		}
// 		wg.Done()
// 	}
// 	wg.Add(len(chs))
// 	for _, c := range chs {
// 		go output(c)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

// func doIt(ctx context.Context, mutex *sync.Mutex, url string, search string, mp map[string]int) string {

// 	out := ""

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(nil))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	resp, err := http.DefaultClient.Do(req)

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	mutex.Lock()
// 	mp[url] = strings.Count(string(body), search)
// 	mutex.Unlock()
// 	count := strings.Count(string(body), search)
// 	num := strconv.Itoa(count)
// 	out = "<" + url + ">" + "-" + "<" + num + ">"

// 	return out
// }
