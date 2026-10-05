package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ch1 <- "Waheed"
		wg.Done()
	}()

	go func() {
		wg.Wait()
		ch2 <- "Abouelfethouh"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
