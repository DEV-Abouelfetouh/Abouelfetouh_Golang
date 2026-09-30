package main

import (
	"fmt"
	"sync"
)

func main() {
	const numChannels = 20

	channels := make([]chan string, numChannels)
	for i := range channels {
		channels[i] = make(chan string)
	}

	var wg sync.WaitGroup
	wg.Add(len(channels))

	processChannel := func(ch <-chan string) {
		defer wg.Done()
		for msg := range ch {
			fmt.Println(msg)
		}
	}

	for _, ch := range channels {
		go processChannel(ch)
	}

	for i, ch := range channels {
		go func(id int, c chan string) {
			c <- fmt.Sprintf("Welcome From Chan %d", id+1)
			close(c)
		}(i, ch)
	}

	wg.Wait()
}
