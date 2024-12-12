package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// it says order should be preserved, but the rules don't actually care about order
func blink(stones []int) []int {
	newStones := []int{}
	for i, stone := range stones {
		if stone == 0 {
			stones[i] = 1
		} else if len(strconv.Itoa(stone))%2 == 0 {
			stoneStr := strconv.Itoa(stone)
			midpoint := len(stoneStr) / 2

			left, err := strconv.Atoi(stoneStr[:midpoint])
			if err != nil {
				panic(err)
			}

			right, err := strconv.Atoi(stoneStr[midpoint:])
			if err != nil {
				panic(err)
			}

			stones[i] = left
			newStones = append(newStones, right)
		} else {
			stones[i] = stone * 2024
		}
	}

	stones = append(stones, newStones...)

	return stones
}

func SolvePart1() {
	file, err := os.Open("input/day11input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	stones := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		for _, num := range strings.Split(line, " ") {
			stone, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			stones = append(stones, stone)
		}
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println(len(stones))
}
