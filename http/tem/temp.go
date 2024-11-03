package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func getWeather() int {
    time.Sleep(1 * time.Second)
    return rand.Intn(70) - 30
}

func main() {
    http.HandleFunc("/weather/highload", func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		ch := make(chan int, 1)
		defer close(ch)

		go func() {
			ch <- getWeather()
		}()

		select {
		case v := <-ch:
			_, _ = resp.Write([]byte(fmt.Sprintf("highload - %d", v)))
		case <-ctx.Done():
			http.Error(resp, "failed timeout", http.StatusRequestTimeout)
		}
    })
}