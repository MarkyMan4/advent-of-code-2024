package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// find the index of the first value from vals found in arr
// if no values from val exist in arr, return -1
func findFirstContaining(arr []int, vals []int) int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(vals); j++ {
			if arr[i] == vals[j] {
				return i
			}
		}
	}

	return -1
}

func correctOrder(pageNums []int, rules map[int][]int) []int {
	corrected := make([]int, len(pageNums))
	_ = copy(corrected, pageNums)

	// keep swapping values that are in the wrong order until the order is correct
	for !isOrderCorrect(corrected, rules) {
		for i, page := range corrected {
			if valsAfter, ok := rules[page]; ok {
				incorrectIdx := findFirstContaining(corrected[:i], valsAfter)
				if incorrectIdx != -1 {
					// if it didn't return -1, swap current val with the index of correct val
					temp := corrected[incorrectIdx]
					corrected[incorrectIdx] = page
					corrected[i] = temp
				}
			}
		}
	}

	return corrected
}

func SolvePart2() {
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

		if !isOrderCorrect(pageNums, rules) {
			corrected := correctOrder(pageNums, rules)
			sum += corrected[len(corrected)/2]
		}
	}

	fmt.Println(sum)
}
