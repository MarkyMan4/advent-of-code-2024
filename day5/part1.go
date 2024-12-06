package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseRule(ruleText string) (int, int) {
	rule := strings.Split(ruleText, "|")
	key, err := strconv.Atoi(rule[0])
	if err != nil {
		log.Fatal(err)
	}

	val, err := strconv.Atoi(rule[1])
	if err != nil {
		log.Fatal(err)
	}

	return key, val
}

func parseUpdate(updateText string) []int {
	data := strings.Split(updateText, ",")
	pageNums := []int{}
	for _, item := range data {
		val, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		pageNums = append(pageNums, val)
	}

	return pageNums
}

func sliceContainsAny(arr []int, vals []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(vals); j++ {
			if arr[i] == vals[j] {
				return true
			}
		}
	}

	return false
}

func SolvePart1() {
	file, err := os.Open("input/day5input.txt")
	if err != nil {
		log.Fatal(err)
	}

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
Scan:
	for scanner.Scan() {
		line := scanner.Text()
		pageNums := parseUpdate(line)

		// check if any rules are violated
		// if we make it through this loop, no rules were broken, so take the middle val
		// and add it to the sum
		for i, page := range pageNums {
			if valsAfter, ok := rules[page]; ok {
				if sliceContainsAny(pageNums[:i], valsAfter) {
					continue Scan
				}
			}
		}

		sum += pageNums[len(pageNums)/2]
	}

	fmt.Println(sum)
}
