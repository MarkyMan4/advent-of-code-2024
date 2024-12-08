package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolvePart1() {
	file, err := os.Open("input/day7input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	equations := []equation{}

	for scanner.Scan() {
		line := scanner.Text()
		testVal, operands := parseEquation(line)
		equations = append(equations, equation{testVal, operands})
	}

	chars := []rune{plus, mul}
	sum := 0

	for _, eq := range equations {
		operatorOptions := generateOperandPermutations(len(eq.operands)-1, chars)

		for _, opList := range operatorOptions {
			if evaulate(eq.operands, opList) == eq.testVal {
				sum += eq.testVal
				break
			}
		}
	}

	fmt.Println(sum)
}
