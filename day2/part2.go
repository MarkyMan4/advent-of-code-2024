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
check an individual diff to see if it meets the safe criteria
*/
func isDiffUnsafe(diff int, isPositive bool) bool {
	return (isPositive && diff < 0) || (!isPositive && diff >= 0) || math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3
}

// make a copy of a report, except for one index
func getRptWithoutIdx(report []int, idx int) []int {
	newRpt := []int{}
	for i, val := range report {
		if i != idx {
			newRpt = append(newRpt, val)
		}
	}

	return newRpt
}

/*
compute difference between each val in the list
then make sure each difference is at least one and at most three
and make sure each difference as the same sign to ensure all increasing or all decreasing
*/
func isReportSafeWithDampener(report []int) bool {
	diffs := []int{}
	for i := 0; i < len(report)-1; i++ {
		diffs = append(diffs, report[i+1]-report[i])
	}

	// determine sign by checking which direction the majority of diffs are
	var isPositive bool
	posCount := 0
	negCount := 0
	for _, diff := range diffs {
		if diff >= 0 {
			posCount++
		} else {
			negCount++
		}
	}

	isPositive = posCount > negCount

	badLevels := []int{}
	for i, diff := range diffs {
		if isDiffUnsafe(diff, isPositive) {
			// if not safe, keep track of levels we can try removing from the report
			badLevels = append(badLevels, i, i+1)
		}
	}

	if len(badLevels) == 0 {
		return true
	}

	for _, levelIdx := range badLevels {
		modifiedReport := getRptWithoutIdx(report, levelIdx)

		// using original isReportSafe function from part 1
		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func SolvePart2() {
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

		if isReportSafeWithDampener(report) {
			safeReportCount++
		}
	}

	fmt.Println(safeReportCount)
}
