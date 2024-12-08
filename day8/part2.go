package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
given two antennas of the same frequency, find the antinode positions

for part 2, don't stop finding antinodes after two, keep going down the line until it runs off the map
also include the antennas in antinodes
*/
func findAntinodesCorrected(antenna1 coord, antenna2 coord, w int, h int) []coord {
	a1Rise := antenna1.row - antenna2.row
	a1Run := antenna1.col - antenna2.col
	a2Rise := a1Rise * -1
	a2Run := a1Run * -1

	antinodes := []coord{antenna1, antenna2}
	node1Pos := coord{antenna1.row + a1Rise, antenna1.col + a1Run}
	node2Pos := coord{antenna2.row + a2Rise, antenna2.col + a2Run}

	i := 2
	for isCoordOnMap(node1Pos, w, h) || isCoordOnMap(node2Pos, w, h) {
		if isCoordOnMap(node1Pos, w, h) {
			antinodes = append(antinodes, node1Pos)
		}

		if isCoordOnMap(node2Pos, w, h) {
			antinodes = append(antinodes, node2Pos)
		}

		node1Pos = coord{antenna1.row + (a1Rise * i), antenna1.col + (a1Run * i)}
		node2Pos = coord{antenna2.row + (a2Rise * i), antenna2.col + (a2Run * i)}

		i++
	}

	return antinodes
}

/*
find all antinode positions within the bounds of the map

for part 2, same thing but keep searching for antinodes in line with antennas as long as they're on the map
*/
func computeAllAntinodePositionsCorrected(antennaPositions map[rune][]coord, mapWidth int, mapHeight int) []coord {
	antinodes := []coord{}

	for _, coords := range antennaPositions {
		// for each frequency, compare each antenna position with all others of the same frequency, finding antinodes
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				if i == j {
					continue
				}
				antinodesForPair := findAntinodesCorrected(coords[i], coords[j], mapWidth, mapHeight)

				for _, node := range antinodesForPair {
					if !coordSliceContains(antinodes, node) {
						antinodes = append(antinodes, node)
					}
				}
			}
		}
	}

	return antinodes
}

func SolvePart2() {
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
	antinodes := computeAllAntinodePositionsCorrected(positions, len(antennaMap[0]), len(antennaMap))

	fmt.Println(len(antinodes))
}
