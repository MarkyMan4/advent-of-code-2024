package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
compute difference between each val in the list
then make sure each difference is at least one and at most three
and make sure each difference as the same sign to ensure all increasing or all decreasing
*/
func isReportSafe(report []int) bool {
	diffs := []int{}
	for i := 0; i < len(report)-1; i++ {
		diffs = append(diffs, report[i+1]-report[i])
	}

	isPositive := true
	for i, diff := range diffs {
		if i == 0 && diff < 0 {
			isPositive = false
		}

		// check sign
		if (isPositive && diff < 0) || (!isPositive && diff >= 0) {
			return false
		}

		// check difference
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
	}

	return true
}

func SolvePart1() {
	file, err := os.Open("input/day2input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	safeReportCount := 0

	for scanner.Scan() {
		// process line by line, parse the report and determine if the numbers are safe
		line := scanner.Text()
		reportText := strings.Split(line, " ")
		report := []int{}

		for _, val := range reportText {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, num)
		}

		if isReportSafe(report) {
			safeReportCount++
		}
	}

	fmt.Println(safeReportCount)
}
