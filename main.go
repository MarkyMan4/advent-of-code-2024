package main

import (
	"log"
	"os"

	"github.com/MarkyMan4/advent-of-code-2024/day1"
	"github.com/MarkyMan4/advent-of-code-2024/day2"
	"github.com/MarkyMan4/advent-of-code-2024/day3"
	"github.com/MarkyMan4/advent-of-code-2024/day4"
	"github.com/MarkyMan4/advent-of-code-2024/day5"
	"github.com/MarkyMan4/advent-of-code-2024/day6"
	"github.com/MarkyMan4/advent-of-code-2024/day7"
	"github.com/MarkyMan4/advent-of-code-2024/day8"
	"github.com/MarkyMan4/advent-of-code-2024/day9"
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
	case "21":
		day2.SolvePart1()
	case "22":
		day2.SolvePart2()
	case "31":
		day3.SolvePart1()
	case "32":
		day3.SolvePart2()
	case "41":
		day4.SolvePart1()
	case "42":
		day4.SolvePart2()
	case "51":
		day5.SolvePart1()
	case "52":
		day5.SolvePart2()
	case "61":
		day6.SolvePart1()
	case "62":
		day6.SolvePart2()
	case "71":
		day7.SolvePart1()
	case "72":
		day7.SolvePart2()
	case "81":
		day8.SolvePart1()
	case "82":
		day8.SolvePart2()
	case "91":
		day9.SolvePart1()
	case "92":
		day9.SolvePart2()
	}
}
