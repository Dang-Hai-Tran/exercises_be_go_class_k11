package main

import "fmt"

func main() {
	var nam Person = CreatePerson("Nam", "Construction Engineer", 1993)
	fmt.Println(`Nam's suitability job:` , nam.checkJobSuitability())

	var str = `Hello World`
	var m = CreateMapByString(str)
	fmt.Println(`Map created by string`, str, `:`, m)

	var slicePerson = ReadFile("exec02/a.txt")
	fmt.Println("Persons from file a.txt: ", slicePerson)
}
