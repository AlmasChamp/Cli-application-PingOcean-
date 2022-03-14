package main

import (
	"log"
	"os"
	"parser/service"
)

func main() {

	args := os.Args[1:]

	if err := service.Start(args); err != nil {
		log.Fatal(err)
	}

}

// FAN-IN FAN-OUT

// func main() {

// 	args := os.Args[1:]

// 	urls := strings.Split(args[1], ",")
// 	search := args[3]
// 	mutex := new(sync.Mutex)
// 	mp := make(map[string]int)

// 	input := gen(urls)
// 	ch1 := info(mutex, search, mp, input)
// 	ch2 := info(mutex, search, mp, input)
// 	ch3 := info(mutex, search, mp, input)

// 	for n := range merge(ch1, ch2, ch3) {
// 		fmt.Println(n)
// 	}
// 	fmt.Println("End")

// }

// func gen(urls []string) <-chan string {
// 	out := make(chan string)

// 	go func() {
// 		for _, url := range urls {
// 			out <- url
// 		}
// 		close(out)
// 	}()

// 	return out
// }

// func info(mutex *sync.Mutex, search string, mp map[string]int, input <-chan string) <-chan string {

// 	ctx := context.Background()
// 	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

// 	out := make(chan string)
// 	go func() {
// 		for url := range input {
// 			out <- doIt(ctx, mutex, url, search, mp)
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func merge(chs ...<-chan string) <-chan string {
// 	var wg sync.WaitGroup

// 	out := make(chan string)

// 	output := func(ch <-chan string) {
// 		for url := range ch {
// 			out <- url
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
