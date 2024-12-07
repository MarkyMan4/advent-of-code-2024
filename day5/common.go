package day5

import (
	"log"
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

func isOrderCorrect(pageNums []int, rules map[int][]int) bool {
	for i, page := range pageNums {
		if valsAfter, ok := rules[page]; ok {
			if sliceContainsAny(pageNums[:i], valsAfter) {
				return false
			}
		}
	}

	return true
}
