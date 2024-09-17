package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Printf("Employees: %v\n", GetEmployees())
	employees := GetEmployees()
	start := time.Now()
	numJobs := len(employees)
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	for w := 1; w <= 2; w++ {
		go Worker(w, jobs, results, employees)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time execution: %s\n", elapsed)
}
