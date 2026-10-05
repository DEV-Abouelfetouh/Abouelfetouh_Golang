package main

import (
	"fmt"
	"sort"
)

type Task struct {
	ID       int
	Name     string
	Payload  map[string]any
	Priority int
}

type ByPriority []Task

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority > a[j].Priority }

func main() {
	tasks := []Task{
		{
			ID:       101,
			Name:     "Send Welcome Email",
			Payload:  map[string]any{"user_id": 55, "template": "welcome_v2"},
			Priority: 2,
		},
		{
			ID:       102,
			Name:     "Process Payment",
			Payload:  map[string]any{"amount": 250.75, "currency": "USD"},
			Priority: 5,
		},
		{
			ID:       103,
			Name:     "Cleanup Logs",
			Payload:  map[string]any{"older_than_days": 30},
			Priority: 1,
		},
	}

	sort.Sort(ByPriority(tasks))

	fmt.Println("Tasks sorted by Priority (High to Low):")
	for _, t := range tasks {
		fmt.Printf("ID: %d | Name: %-20s | Priority: %d | Payload: %v\n", t.ID, t.Name, t.Priority, t.Payload)
	}
}
