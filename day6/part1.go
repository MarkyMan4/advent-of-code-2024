package day6

import "fmt"

type coord struct {
	row int
	col int
}

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

func coordSliceContains(cs []coord, c coord) bool {
	for i := range cs {
		if cs[i] == c {
			return true
		}
	}

	return false
}

func findGuardPosition(worldMap []string) coord {
	for i, line := range worldMap {
		for j, tile := range line {
			if tile == '^' {
				return coord{i, j}
			}
		}
	}

	// should never reach this if map is correct
	return coord{0, 0}
}

func SolvePart1() {
	worldMap := readMap()
	guardPos := findGuardPosition(worldMap)
	guardDir := up
	coordsVisited := []coord{}

WalkLoop:
	for {
		if !coordSliceContains(coordsVisited, guardPos) {
			coordsVisited = append(coordsVisited, guardPos)
		}

		// make the guard take one step
		switch guardDir {
		case up:
			if guardPos.row == 0 {
				break WalkLoop
			} else if worldMap[guardPos.row-1][guardPos.col] == '#' {
				guardDir = right
			} else {
				guardPos.row--
			}
		case right:
			if guardPos.col == len(worldMap[0])-1 {
				break WalkLoop
			} else if worldMap[guardPos.row][guardPos.col+1] == '#' {
				guardDir = down
			} else {
				guardPos.col++
			}
		case down:
			if guardPos.row == len(worldMap)-1 {
				break WalkLoop
			} else if worldMap[guardPos.row+1][guardPos.col] == '#' {
				guardDir = left
			} else {
				guardPos.row++
			}
		case left:
			if guardPos.col == 0 {
				break WalkLoop
			} else if worldMap[guardPos.row][guardPos.col-1] == '#' {
				guardDir = up
			} else {
				guardPos.col--
			}
		}
	}

	fmt.Println(len(coordsVisited))
}
