package day10

import "fmt"

func findAllTrailheadPaths(trailMap [][]int) int {
	scoreSum := 0
	for i := range trailMap {
		for j := range trailMap[i] {
			if trailMap[i][j] == 0 {
				scoreSum += findPathsFromTrailhead(coord{i, j}, trailMap)
			}
		}
	}

	return scoreSum
}

// same thing as part one, but don't find unique
func findPathsFromTrailhead(c coord, trailMap [][]int) int {
	endpoints := findEndpointsReached(c, trailMap, []coord{})
	return len(endpoints)
}

func SolvePart2() {
	trailMap := readTrailMap()
	fmt.Println(findAllTrailheadPaths(trailMap))
}
