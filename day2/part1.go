package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
