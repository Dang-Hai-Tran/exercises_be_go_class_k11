package main

type Person struct {
	name      string
	job       string
	birthYear int
}

func (p Person) checkJobSuitability() bool {
	if p.birthYear%len(p.name) == 0 {
		return true
	}
	return false
}

func CreatePerson(name string, job string, birthYear int) Person {
	var person = Person{name: name, job: job, birthYear: birthYear}
	return person
}
