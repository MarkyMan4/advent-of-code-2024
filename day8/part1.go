package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
given two antennas of the same frequency, find the antinode positions
for two nodes, we will always have two antinodes
*/
func findAntinodes(antenna1 coord, antenna2 coord) (coord, coord) {
	a1Rise := antenna1.row - antenna2.row
	a1Run := antenna1.col - antenna2.col
	a2Rise := a1Rise * -1
	a2Run := a1Run * -1

	node1Pos := coord{antenna1.row + a1Rise, antenna1.col + a1Run}
	node2Pos := coord{antenna2.row + a2Rise, antenna2.col + a2Run}

	return node1Pos, node2Pos
}

/*
find all antinode positions within the bounds of the map
*/
func computeAllAntinodePositions(antennaPositions map[rune][]coord, mapWidth int, mapHeight int) []coord {
	antinodes := []coord{}

	for _, coords := range antennaPositions {
		// for each frequency, compare each antenna position with all others of the same frequency, finding antinodes
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				if i == j {
					continue
				}
				antinode1, antinode2 := findAntinodes(coords[i], coords[j])

				// save position if the antinode is on the map AND if we don't already have this position saved
				if isCoordOnMap(antinode1, mapWidth, mapHeight) && !coordSliceContains(antinodes, antinode1) {
					antinodes = append(antinodes, antinode1)
				}
				if isCoordOnMap(antinode2, mapWidth, mapHeight) && !coordSliceContains(antinodes, antinode2) {
					antinodes = append(antinodes, antinode2)
				}
			}
		}
	}

	return antinodes
}

func SolvePart1() {
	file, err := os.Open("input/day8input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	antennaMap := []string{}

	for scanner.Scan() {
		antennaMap = append(antennaMap, scanner.Text())
	}

	positions := findAntennaPositions(antennaMap)
	antinodes := computeAllAntinodePositions(positions, len(antennaMap[0]), len(antennaMap))

	fmt.Println(len(antinodes))
}
