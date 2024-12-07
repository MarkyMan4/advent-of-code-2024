package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolvePart1() {
	file, err := os.Open("input/day5input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	// create a mapping from page numbers to what they are less than
	rules := map[int][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		key, val := parseRule(line)
		rules[key] = append(rules[key], val)
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		pageNums := parseUpdate(line)

		if isOrderCorrect(pageNums, rules) {
			sum += pageNums[len(pageNums)/2]
		}
	}

	fmt.Println(sum)
}
