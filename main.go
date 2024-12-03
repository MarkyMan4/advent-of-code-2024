package main

import (
	"log"
	"os"

	"github.com/MarkyMan4/advent-of-code-2024/day1"
)

/*
provide command line argument to tell the day and part to solve

e.g.
to solve day 3 part 2: go run main.go 32
day 10 part 1: go run main.go 101
etc
*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide an argument telling which day and part to solve")
	}

	dayAndPart := os.Args[1]

	switch dayAndPart {
	case "11":
		day1.SolvePart1()
	case "12":
		day1.SolvePart2()
	}
}
