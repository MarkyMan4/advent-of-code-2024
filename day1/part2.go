package day1

import "fmt"

/*
compute how many time each distinct number appears in the list and return a map
where the keys are numbers from the list and the value is how many times that number
appeared
*/
func computeFrequencies(vals []int) map[int]int {
	freqs := make(map[int]int)
	for _, val := range vals {
		freqs[val]++
	}

	return freqs
}

func SolvePart2() {
	left, right := readLists()
	freqs := computeFrequencies(right)
	sum := 0

	for _, val := range left {
		sum += val * freqs[val]
	}

	fmt.Println(sum)
}
