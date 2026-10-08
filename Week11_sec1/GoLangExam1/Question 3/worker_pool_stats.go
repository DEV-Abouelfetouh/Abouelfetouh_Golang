package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
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

type TaskResult struct {
	TaskID   int
	Status   string
	Retries  int
	Duration time.Duration
	Err      error
}

type PoolStats struct {
	Processed int64
	Succeeded int64
	Failed    int64
	Panicked  int64
	Retried   int64
}

type Handler interface {
	Handle(ctx context.Context, t AppTask) error
}

type EmailHandler struct{}

func (h *EmailHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[EmailHandler] Processing task ID %d\n", t.ID)
	if t.ID == 1 {
		panic("unexpected nil pointer dereference in email template renderer")
	}
	if t.ID == 4 {
		return &RetryableError{Err: errors.New("temporary SMTP connection timeout")}
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}

type ResizeImageHandler struct{}

func (h *ResizeImageHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ResizeImageHandler] Processing task ID %d\n", t.ID)
	time.Sleep(150 * time.Millisecond)
	return nil
}

type ReportHandler struct{}

func (h *ReportHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ReportHandler] Processing task ID %d\n", t.ID)
	time.Sleep(200 * time.Millisecond)
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

func (d *TaskDispatcher) Dispatch(ctx context.Context, t AppTask) (res TaskResult) {
	start := time.Now()
	res.TaskID = t.ID

	handler, exists := d.handlers[t.Type]
	if !exists {
		res.Status = "failed"
		res.Duration = time.Since(start)
		res.Err = fmt.Errorf("no handler registered for task type: %s", t.Type)
		return res
	}

	var panicked bool
	var panicVal interface{}

	defer func() {
		if r := recover(); r != nil {
			panicked = true
			panicVal = r
			res.Status = "panicked"
			res.Duration = time.Since(start)
			res.Err = fmt.Errorf("panic recovered: %v", panicVal)
		}
	}()

	maxRetries := 3
	var handlerErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		handlerErr = handler.Handle(ctx, t)
		if handlerErr == nil {
			res.Status = "succeeded"
			res.Duration = time.Since(start)
			return res
		}

		var retryErr *RetryableError
		if !errors.As(handlerErr, &retryErr) {
			res.Status = "failed"
			res.Retries = attempt - 1
			res.Duration = time.Since(start)
			res.Err = fmt.Errorf("non-retryable error: %w", handlerErr)
			return res
		}

		res.Retries = attempt
		log.Printf("[Warning] Task %d failed with retryable error (Attempt %d/%d): %v", t.ID, attempt, maxRetries, handlerErr)

		if attempt == maxRetries {
			break
		}

		select {
		case <-ctx.Done():
			res.Status = "failed"
			res.Duration = time.Since(start)
			res.Err = ctx.Err()
			return res
		case <-time.After(time.Duration(attempt) * 100 * time.Millisecond):
		}
	}

	if panicked {
		return res
	}

	res.Status = "failed"
	res.Duration = time.Since(start)
	res.Err = fmt.Errorf("failed after %d max retries: %w", maxRetries, handlerErr)
	return res
}

type WorkerPool struct {
	numWorkers int
	dispatcher *TaskDispatcher
	taskChan   chan AppTask
	Stats      PoolStats
	resultsMu  sync.Mutex
	results    []TaskResult
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int, dispatcher *TaskDispatcher, bufferSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		dispatcher: dispatcher,
		taskChan:   make(chan AppTask, bufferSize),
		results:    make([]TaskResult, 0),
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
					res := pool.dispatcher.Dispatch(ctx, task)

					// Update atomic shared statistics safely
					atomic.AddInt64(&pool.Stats.Processed, 1)
					if res.Retries > 0 {
						atomic.AddInt64(&pool.Stats.Retried, 1)
					}
					switch res.Status {
					case "succeeded":
						atomic.AddInt64(&pool.Stats.Succeeded, 1)
					case "failed":
						atomic.AddInt64(&pool.Stats.Failed, 1)
					case "panicked":
						atomic.AddInt64(&pool.Stats.Panicked, 1)
					}

					pool.resultsMu.Lock()
					pool.results = append(pool.results, res)
					pool.resultsMu.Unlock()
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

func (p *WorkerPool) PrintSummary() {
	p.resultsMu.Lock()
	defer p.resultsMu.Unlock()

	var totalDuration time.Duration
	for _, res := range p.results {
		totalDuration += res.Duration
	}

	var avgDuration time.Duration
	if len(p.results) > 0 {
		avgDuration = totalDuration / time.Duration(len(p.results))
	}

	fmt.Println("\n========== ATOMIC SHARED STATS SUMMARY ==========")
	fmt.Printf("Processed (Atomic) : %d\n", atomic.LoadInt64(&p.Stats.Processed))
	fmt.Printf("Succeeded (Atomic) : %d\n", atomic.LoadInt64(&p.Stats.Succeeded))
	fmt.Printf("Failed (Atomic)    : %d\n", atomic.LoadInt64(&p.Stats.Failed))
	fmt.Printf("Panicked (Atomic)  : %d\n", atomic.LoadInt64(&p.Stats.Panicked))
	fmt.Printf("Retried (Atomic)   : %d\n", atomic.LoadInt64(&p.Stats.Retried))
	fmt.Printf("Average Duration   : %v\n", avgDuration)
	fmt.Println("=================================================")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dispatcher := NewTaskDispatcher()
	pool := NewWorkerPool(3, dispatcher, 10)
	Start(ctx, pool)

	go func() {
		tasks := []AppTask{
			{ID: 1, Type: "email"},
			{ID: 2, Type: "resize_image"},
			{ID: 3, Type: "report"},
			{ID: 4, Type: "email"},
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
	pool.PrintSummary()
}
