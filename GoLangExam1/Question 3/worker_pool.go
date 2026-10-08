package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

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
	select {
	case <-time.After(500 * time.Millisecond):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

type ResizeImageHandler struct{}

func (h *ResizeImageHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ResizeImageHandler] Processing task ID %d\n", t.ID)
	select {
	case <-time.After(500 * time.Millisecond):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

type ReportHandler struct{}

func (h *ReportHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Printf("[ReportHandler] Processing task ID %d\n", t.ID)
	select {
	case <-time.After(500 * time.Millisecond):
	case <-ctx.Done():
		return ctx.Err()
	}
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

func (d *TaskDispatcher) Dispatch(ctx context.Context, t AppTask) error {
	handler, exists := d.handlers[t.Type]
	if !exists {
		return fmt.Errorf("no handler registered for task type: %s", t.Type)
	}
	return handler.Handle(ctx, t)
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
					fmt.Printf("Worker %d shutting down gracefully...\n", workerID)
					return
				case task, ok := <-pool.taskChan:
					if !ok {
						fmt.Printf("Worker %d task channel closed, exiting...\n", workerID)
						return
					}
					if err := pool.dispatcher.Dispatch(ctx, task); err != nil {
						log.Printf("Worker %d error processing task %d: %v", workerID, task.ID, err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dispatcher := NewTaskDispatcher()

	pool := NewWorkerPool(3, dispatcher, 10)
	Start(ctx, pool)

	go func() {
		taskTypes := []string{"email", "resize_image", "report"}
		for i := 1; i <= 6; i++ {
			task := AppTask{
				ID:   i,
				Type: taskTypes[(i-1)%len(taskTypes)],
			}
			if err := pool.Submit(ctx, task); err != nil {
				log.Printf("Failed to submit task %d: %v", i, err)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	pool.Stop()
	fmt.Println("Worker pool stopped successfully.")
}
