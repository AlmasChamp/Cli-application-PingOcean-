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
	ctx,_ = context.WithTimeout(ctx,10*time.Second)

	mutex:= new(sync.Mutex)
	wg:= &sync.WaitGroup{}
	ch := make(chan string,3)

	for i := 0; i < len(urls); i++{
		wg.Add(1)
		go doIt(ctx,ch,wg,mutex,mp,search) 
	}

	for _, url := range urls {
		ch <- url
	}
	
	close(ch)
	wg.Wait()
	fmt.Println("End")

}

func doIt(ctx context.Context,ch chan string,wg *sync.WaitGroup, mutex *sync.Mutex,mp map[string]int,search string){
	
	url := <-ch
	
	// resp,err := http.Get(url)
	
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
	fmt.Println("<",url,">","-","<",strings.Count(string(body), search),">")

	wg.Done()
}
