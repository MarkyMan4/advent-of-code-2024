package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
take a line from the file and return the two integers
*/
func splitLine(line string) (int, int) {
	nums := strings.Split(line, "   ")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	return num1, num2
}

func readLists() ([]int, []int) {
	file, err := os.Open("input/day1input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		num1, num2 := splitLine(line)
		left = append(left, num1)
		right = append(right, num2)
	}

	return left, right
}
