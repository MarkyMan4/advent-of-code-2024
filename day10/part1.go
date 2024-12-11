package day10

import (
	"fmt"
)

func findTrailheadScores(trailMap [][]int) int {
	scoreSum := 0
	for i := range trailMap {
		for j := range trailMap[i] {
			if trailMap[i][j] == 0 {
				scoreSum += scoreTrailhead(coord{i, j}, trailMap)
			}
		}
	}

	return scoreSum
}

func scoreTrailhead(c coord, trailMap [][]int) int {
	endpoints := findEndpointsReached(c, trailMap, []coord{})
	uniqueEndpoints := []coord{}
	for _, e := range endpoints {
		if !coordSliceContains(uniqueEndpoints, e) {
			uniqueEndpoints = append(uniqueEndpoints, e)
		}
	}

	return len(uniqueEndpoints)
}

func SolvePart1() {
	trailMap := readTrailMap()
	fmt.Println(findTrailheadScores(trailMap))
}
