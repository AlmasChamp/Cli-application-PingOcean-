package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

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
