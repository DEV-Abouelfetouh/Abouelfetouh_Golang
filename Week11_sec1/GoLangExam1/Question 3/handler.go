package main

import (
	"context"
	"fmt"
	"log"
)

type AppTask struct {
	Type    string
	Payload map[string]interface{}
}

type Handler interface {
	Handle(ctx context.Context, t AppTask) error
}

type EmailHandler struct{}

func (h *EmailHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Println("Processing and sending email...")
	return nil
}

type ResizeImageHandler struct{}

func (h *ResizeImageHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Println("Processing image resize...")
	return nil
}

type ReportHandler struct{}

func (h *ReportHandler) Handle(ctx context.Context, t AppTask) error {
	fmt.Println("Generating report...")
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

func main() {
	ctx := context.Background()
	dispatcher := NewTaskDispatcher()

	tasks := []AppTask{
		{Type: "email"},
		{Type: "resize_image"},
		{Type: "report"},
	}

	for _, task := range tasks {
		if err := dispatcher.Dispatch(ctx, task); err != nil {
			log.Printf("Error processing task %s: %v", task.Type, err)
		}
	}
}
