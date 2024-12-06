package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolvePart1() {
	file, err := os.Open("input/day4input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	xmasCount := 0
	for i, line := range lines {
		for j, c := range line {
			if c != 'X' {
				continue
			}

			oneUp := i - 1
			twoUp := i - 2
			threeUp := i - 3
			oneRight := j + 1
			twoRight := j + 2
			threeRight := j + 3
			oneDown := i + 1
			twoDown := i + 2
			threeDown := i + 3
			oneLeft := j - 1
			twoLeft := j - 2
			threeLeft := j - 3

			if threeUp >= 0 && lines[oneUp][j] == 'M' && lines[twoUp][j] == 'A' && lines[threeUp][j] == 'S' {
				xmasCount++
			}
			if threeUp >= 0 && threeRight < len(line) && lines[oneUp][oneRight] == 'M' && lines[twoUp][twoRight] == 'A' && lines[threeUp][threeRight] == 'S' {
				xmasCount++
			}
			if threeRight < len(line) && line[oneRight] == 'M' && line[twoRight] == 'A' && line[threeRight] == 'S' {
				xmasCount++
			}
			if threeDown < len(lines) && threeRight < len(line) && lines[oneDown][oneRight] == 'M' && lines[twoDown][twoRight] == 'A' && lines[threeDown][threeRight] == 'S' {
				xmasCount++
			}
			if threeDown < len(lines) && lines[oneDown][j] == 'M' && lines[twoDown][j] == 'A' && lines[threeDown][j] == 'S' {
				xmasCount++
			}
			if threeDown < len(lines) && threeLeft >= 0 && lines[oneDown][oneLeft] == 'M' && lines[twoDown][twoLeft] == 'A' && lines[threeDown][threeLeft] == 'S' {
				xmasCount++
			}
			if threeLeft >= 0 && line[oneLeft] == 'M' && line[twoLeft] == 'A' && line[threeLeft] == 'S' {
				xmasCount++
			}
			if threeLeft >= 0 && threeUp >= 0 && lines[oneUp][oneLeft] == 'M' && lines[twoUp][twoLeft] == 'A' && lines[threeUp][threeLeft] == 'S' {
				xmasCount++
			}
		}
	}

	fmt.Println(xmasCount)
}
