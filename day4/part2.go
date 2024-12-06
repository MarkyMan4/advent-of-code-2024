package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolvePart2() {
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
			if c != 'A' {
				continue
			}

			up := i - 1
			right := j + 1
			down := i + 1
			left := j - 1

			// if any directions out of range, skip
			if up < 0 || left < 0 || down >= len(lines) || right >= len(line) {
				continue
			}

			diagonalOne := string(lines[up][left]) + string(c) + string(lines[down][right])
			diagonalTwo := string(lines[up][right]) + string(c) + string(lines[down][left])

			if (diagonalOne == "MAS" || diagonalOne == "SAM") && (diagonalTwo == "MAS" || diagonalTwo == "SAM") {
				xmasCount++
			}
		}
	}

	fmt.Println(xmasCount)
}
