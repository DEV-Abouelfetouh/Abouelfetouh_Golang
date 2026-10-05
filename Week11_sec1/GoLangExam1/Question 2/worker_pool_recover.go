package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
}

func process(job Job) int {
	if job.ID%3 == 0 {
		panic(fmt.Sprintf("simulated crash on job %d", job.ID))
	}
	return job.ID * 10
}

func worker(id int, jobs <-chan Job, results chan<- Result, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		func() {
			defer func() {
				if r := recover(); r != nil {
					errs <- fmt.Errorf("worker %d recovered from panic on job %d: %v", id, job.ID, r)
				}
			}()

			val := process(job)
			results <- Result{JobID: job.ID, Value: val}
		}()
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	errs := make(chan error, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, errs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	wg.Wait()
	close(results)
	close(errs)

	fmt.Println("--- Results ---")
	for res := range results {
		fmt.Printf("Job %d -> Result: %d\n", res.JobID, res.Value)
	}

	fmt.Println("\n--- Errors ---")
	for err := range errs {
		fmt.Println(err)
	}

	fmt.Println("\nProgram finished successfully without crashing!")
}
