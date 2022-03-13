package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

func Start(args []string) error {

	if len(args) < 4 {
		return errors.New("ERROR: not enough arguments. Example: -urls https://habr.com -search script")
	} else if len(args) > 4 {
		return errors.New("ERROR: more than 4 arguments. Example: -urls https://habr.com -search script")
	} else if args[0] != "-urls" {
		return errors.New("ERROR: invalid flag, should be -urls")
	} else if args[2] != "-search" {
		return errors.New("ERROR: invalid flag, should be -search")
	}

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

	return nil
}
