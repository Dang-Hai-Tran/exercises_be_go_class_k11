package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- string, employees []Employee) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- fmt.Sprintf("Employee %d: salary/age = %d", j-1, employees[j-1].EmployeeSalary/employees[j-1].EmployeeAge)
	}
}
