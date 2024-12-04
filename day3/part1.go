package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func SolvePart1() {
	file, err := os.Open("input/day3input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mulRegex, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if err != nil {
		log.Fatal(err)
	}

	numRegex, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := mulRegex.FindAll([]byte(line), -1)
		for _, m := range matches {
			nums := numRegex.FindAll(m, 2)
			x, err := strconv.Atoi(string(nums[0]))
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(string(nums[1]))
			if err != nil {
				log.Fatal(err)
			}

			sum += x * y
		}
	}

	fmt.Println(sum)
}
