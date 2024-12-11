package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

func coordSliceContains(cs []coord, c coord) bool {
	for _, val := range cs {
		if val.row == c.row && val.col == c.col {
			return true
		}
	}

	return false
}

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
	return len(endpoints)
}

func findEndpointsReached(c coord, trailMap [][]int, foundEndpoints []coord) []coord {
	if trailMap[c.row][c.col] == 9 {
		return []coord{{c.row, c.col}}
	}

	endpoints := []coord{}

	// only move to nodes that are not nil and are exactly one level above current node
	if c.row-1 >= 0 && trailMap[c.row-1][c.col]-trailMap[c.row][c.col] == 1 {
		endpointsUp := findEndpointsReached(coord{c.row - 1, c.col}, trailMap, foundEndpoints)
		endpoints = append(endpoints, endpointsUp...)
	}
	if c.col+1 < len(trailMap[c.row]) && trailMap[c.row][c.col+1]-trailMap[c.row][c.col] == 1 {
		endpointsRight := findEndpointsReached(coord{c.row, c.col + 1}, trailMap, foundEndpoints)
		endpoints = append(endpoints, endpointsRight...)
	}
	if c.row+1 < len(trailMap) && trailMap[c.row+1][c.col]-trailMap[c.row][c.col] == 1 {
		endpointsDown := findEndpointsReached(coord{c.row + 1, c.col}, trailMap, foundEndpoints)
		endpoints = append(endpoints, endpointsDown...)
	}
	if c.col-1 >= 0 && trailMap[c.row][c.col-1]-trailMap[c.row][c.col] == 1 {
		endpointsLeft := findEndpointsReached(coord{c.row, c.col - 1}, trailMap, foundEndpoints)
		endpoints = append(endpoints, endpointsLeft...)
	}

	uniqueEndpoints := []coord{}
	for _, e := range endpoints {
		if !coordSliceContains(uniqueEndpoints, e) {
			uniqueEndpoints = append(uniqueEndpoints, e)
		}
	}

	return uniqueEndpoints
}

func SolvePart1() {
	file, err := os.Open("input/day10input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	trailMap := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		tiles := strings.Split(line, "")
		row := []int{}

		for _, t := range tiles {
			val, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}

			row = append(row, val)
		}

		trailMap = append(trailMap, row)
	}

	fmt.Println(findTrailheadScores(trailMap))
}
