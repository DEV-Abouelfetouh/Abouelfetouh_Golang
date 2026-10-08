package main

import (
	"errors"
	"fmt"
	"time"
)

func Retry(times int, fn func() error) error {
	var lastErr error
	backoff := 1 * time.Second

	for i := 0; i < times; i++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if i < times-1 {
			time.Sleep(backoff)
			backoff *= 2
		}
	}

	return fmt.Errorf("after %d attempts: %w", times, lastErr)
}

func main() {
	attemptCount := 0
	failingFunc := func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("connection timeout")
		}
		return nil
	}

	err := Retry(3, failingFunc)
	if err != nil {
		fmt.Println("Failed:", err)
	} else {
		fmt.Println("Success!")
	}

	alwaysFails := func() error {
		return errors.New("server unreachable")
	}

	err2 := Retry(2, alwaysFails)
	if err2 != nil {
		fmt.Println("Failed:", err2)
	}
}
