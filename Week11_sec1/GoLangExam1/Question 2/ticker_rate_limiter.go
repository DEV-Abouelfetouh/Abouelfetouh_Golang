package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started at %s\n", id, time.Now().Format("15:04:05.000"))
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	rate := 5
	totalTasks := 20

	interval := time.Second / time.Duration(rate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for i := 1; i <= totalTasks; i++ {
		<-ticker.C
		wg.Add(1)
		go processTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("All 20 tasks processed through the rate limiter.")
}
