package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []Person {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ret []Person
	for scanner.Scan() {
		line := scanner.Text()
		sliceLine := strings.Split(line, "|")
		if len(sliceLine) != 3 {
			log.Println("Line format incorrect")
			return nil
		}
		name, job, birthYearStr := sliceLine[0], sliceLine[1], sliceLine[2]
		birthYear, err := strconv.Atoi(birthYearStr)
		if err != nil {
			log.Println("Error when convert birth year to int")
			return nil
		}
		person := Person{name, job, birthYear}
		ret = append(ret, person)
	}
	return ret
}
