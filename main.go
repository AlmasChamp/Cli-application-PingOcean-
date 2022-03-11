package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {

	args := os.Args[1:]

	urls := strings.Split(args[1], ",")
	search := args[3]
	mp := make(map[string]int)

	mutex:= new(sync.Mutex)
	wg:= &sync.WaitGroup{}
	ch := make(chan string,3)

	for i := 0; i < len(urls); i++{
		wg.Add(1)
		go doIt(mp,search,ch,wg,mutex) 
	}

	for _, url := range urls {
		ch <- url
	}
	
	close(ch)
	wg.Wait()
	fmt.Println("End")

}

func doIt(mp map[string]int,search string,ch chan string,wg *sync.WaitGroup, mutex *sync.Mutex){
	url := <-ch
	
	mutex.Lock()
	resp,err := http.Get(url)
	
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	mp[url] = strings.Count(string(body), search)
	fmt.Println("<",url,">","-","<",strings.Count(string(body), search),">")
	mutex.Unlock()

	wg.Done()
}
