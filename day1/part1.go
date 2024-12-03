package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func SolvePart1() {
	file, err := os.Open("input/day1input.txt")
	if err != nil {
		log.Fatalln(err)
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

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	totalDiff := 0
	for i := 0; i < len(left); i++ {
		diff := int(math.Abs(float64(left[i]) - float64(right[i])))
		totalDiff += diff
	}

	fmt.Println(totalDiff)
}
