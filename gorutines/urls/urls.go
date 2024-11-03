package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var urls = []string{
		"http://pornhub.com/video/",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
		"http://pornhub.com",
		"http://utmn.ru",
	}
	ch := make(chan string)
	sema := make(chan struct{}, 1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-sema
			}()
			sema <- struct{}{}
			select {
			case <- ctx.Done():
				return
			default:
			}
			res, err := http.Get(url)
			defer res.Body.Close()
			if err != nil || res.StatusCode != http.StatusOK {
				ch <- fmt.Sprintf("%s: not ok", url)
				return
			}
			ch <- fmt.Sprintf("%s: ok", url)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}