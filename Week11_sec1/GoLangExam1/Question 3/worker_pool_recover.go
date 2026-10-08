package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type RetryableError struct {
	Err error
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("retryable error: %v", e.Err)
}

func (e *RetryableError) Unwrap() error {
	return e.Err
}

type AppTask struct {
	ID      int
	Type    string
	Payload map[string]interface{}
}

type Handler interface {
	Handle(ctx context.Context, t AppTask) error
}

type EmailHandler struct{}

func (h *EmailHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[EmailHandler] Processing task ID %d\n", t.ID)
	if t.ID == 1 {
		// Simulate an unexpected runtime panic (e.g., nil pointer dereference)
		panic("unexpected nil pointer dereference in email template renderer")
	}
	return nil
}

type ResizeImageHandler struct{}

func (h *ResizeImageHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ResizeImageHandler] Processing task ID %d\n", t.ID)
	return nil
}

type ReportHandler struct{}

func (h *ReportHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ReportHandler] Processing task ID %d\n", t.ID)
	return nil
}

type TaskDispatcher struct {
	handlers map[string]Handler
}

func NewTaskDispatcher() *TaskDispatcher {
	d := &TaskDispatcher{
		handlers: make(map[string]Handler),
	}
	d.Register("email", &EmailHandler{})
	d.Register("resize_image", &ResizeImageHandler{})
	d.Register("report", &ReportHandler{})
	return d
}

func (d *TaskDispatcher) Register(taskType string, handler Handler) {
	d.handlers[taskType] = handler
}

func (d *TaskDispatcher) Dispatch(ctx context.Context, t AppTask) (err error) {
	handler, exists := d.handlers[t.Type]
	if !exists {
		return fmt.Errorf("no handler registered for task type: %s", t.Type)
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during task execution: %v", r)
		}
	}()

	maxRetries := 3
	var handlerErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		handlerErr = handler.Handle(ctx, t)
		if handlerErr == nil {
			return nil
		}

		var retryErr *RetryableError
		if !errors.As(handlerErr, &retryErr) {
			return fmt.Errorf("non-retryable error on task %d: %w", t.ID, handlerErr)
		}

		log.Printf("[Warning] Task %d failed with retryable error (Attempt %d/%d): %v", t.ID, attempt, maxRetries, handlerErr)

		if attempt == maxRetries {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(attempt) * 200 * time.Millisecond):
		}
	}

	return fmt.Errorf("task %d failed after %d maximum retries: %w", t.ID, maxRetries, handlerErr)
}

type WorkerPool struct {
	numWorkers int
	dispatcher *TaskDispatcher
	taskChan   chan AppTask
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int, dispatcher *TaskDispatcher, bufferSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		dispatcher: dispatcher,
		taskChan:   make(chan AppTask, bufferSize),
	}
}

func Start(ctx context.Context, pool *WorkerPool) {
	for i := 1; i <= pool.numWorkers; i++ {
		pool.wg.Add(1)
		go func(workerID int) {
			defer pool.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case task, ok := <-pool.taskChan:
					if !ok {
						return
					}
					if err := pool.dispatcher.Dispatch(ctx, task); err != nil {
						log.Printf("Worker %d caught error/panic for task %d: %v", workerID, task.ID, err)
					}
				}
			}
		}(i)
	}
}

func (p *WorkerPool) Submit(ctx context.Context, task AppTask) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case p.taskChan <- task:
		return nil
	}
}

func (p *WorkerPool) Stop() {
	close(p.taskChan)
	p.wg.Wait()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dispatcher := NewTaskDispatcher()
	pool := NewWorkerPool(3, dispatcher, 10)
	Start(ctx, pool)

	go func() {
		tasks := []AppTask{
			{ID: 1, Type: "email"}, // Will trigger a panic
			{ID: 2, Type: "resize_image"},
			{ID: 3, Type: "report"},
		}
		for _, task := range tasks {
			if err := pool.Submit(ctx, task); err != nil {
				log.Printf("Failed to submit task %d: %v", task.ID, err)
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)
	pool.Stop()
	fmt.Println("Worker pool stopped successfully.")
}
