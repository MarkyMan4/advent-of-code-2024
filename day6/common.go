package day6

import (
	"bufio"
	"log"
	"os"
)

func readMap() []string {
	file, err := os.Open("input/day6input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
