package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Define a struct to match the structure of the JSON response
type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// Response struct to match the API response structure
type Response struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}

func GetEmployees() []Employee {
	// Perform the GET request
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Define a variable to store the unmarshaled data
	var employeesResponse Response

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &employeesResponse)
	if err != nil {
		panic(err)
	}

	// Print out the unmarshaled data
	// for _, employee := range employeesResponse.Data {
	// 	fmt.Printf("ID: %d, Name: %s, Salary: %d, Age: %d\n",
	// 		employee.ID, employee.EmployeeName, employee.EmployeeSalary, employee.EmployeeAge)
	// }
	return employeesResponse.Data
}
