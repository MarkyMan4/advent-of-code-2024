package day1

import (
	"fmt"
	"math"
	"sort"
)

func SolvePart1() {
	left, right := readLists()

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
