package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	testVal  int
	operands []int
}

// operators
const (
	plus = '+'
	mul  = '*'
)

func parseEquation(text string) (int, []int) {
	parts := strings.Split(text, ":")
	testVal, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	opList := strings.Split(strings.Trim(parts[1], " "), " ")
	operands := []int{}
	for _, val := range opList {
		op, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		operands = append(operands, op)
	}

	return testVal, operands
}

/*
find all permutations of '+' and '*' for a list of length n. This is basically the
same as finding all numbers up to n digits for a base "x" number system, where "x" is len(chars).
In this problem, I know the only operators are + and *, so it's a base 2 system. But this function works for
any base (i.e. it would also work if we included +, -, *, /)

e.g. n = 2
+, +
+, *
*, +
*, *
*/
func generateOperandPermutations(n int, chars []rune) [][]rune {
	permutations := [][]rune{}
	numbers := computeNumbers(n, len(chars))
	for _, num := range numbers {
		opList := []rune{}
		for _, digit := range num {
			opList = append(opList, chars[digit])
		}

		permutations = append(permutations, opList)
	}

	return permutations
}

/*
find all numbers with n digits for a base n numbering system

e.g. digits = 2, base = 3
00
01
02
10
11
12
20
21
22
*/
func computeNumbers(digits int, base int) [][]int {
	numberCount := int(math.Pow(float64(base), float64(digits)))
	nums := [][]int{}

	current := []int{}
	for i := 0; i < digits; i++ {
		current = append(current, 0)
	}

	save := make([]int, len(current))
	copy(save, current)
	nums = append(nums, save)

	// iterate until numberCount - 1 because I've already calculated 0
	for i := 0; i < numberCount-1; i++ {
		for j := range current {
			if current[j] < base-1 {
				current[j]++
				save := make([]int, len(current))
				copy(save, current)
				nums = append(nums, save)
				break

			} else {
				current[j] = 0
			}

		}
	}

	return nums
}

/*
evaulate an expression. Expressions are evaluated left to right, order of operations does not apply
e.g.
operands = [1, 2, 3]
operators = [+, *]

result = (1 + 2) * 3

assumption: len(operators) = len(operands)-1
*/
func evaulate(operands []int, operators []rune) int {
	result := operands[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == plus {
			result += operands[i+1]
		} else if operators[i] == mul {
			result *= operands[i+1]
		}
	}

	return result
}

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
