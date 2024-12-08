package day8

type coord struct {
	row int
	col int
}

func coordSliceContains(coords []coord, val coord) bool {
	for _, c := range coords {
		if c.row == val.row && c.col == val.col {
			return true
		}
	}

	return false
}

/*
look through the map and create a map where keys are antenna frequencies and value is a list of coordinates with that frequency
*/
func findAntennaPositions(antennaMap []string) map[rune][]coord {
	coords := map[rune][]coord{}
	for i := 0; i < len(antennaMap); i++ {
		for j := 0; j < len(antennaMap[i]); j++ {
			char := rune(antennaMap[i][j])
			if antennaMap[i][j] != '.' {
				coords[char] = append(coords[char], coord{i, j})
			}
		}
	}

	return coords
}

func isCoordOnMap(c coord, w int, h int) bool {
	return c.row >= 0 && c.col >= 0 && c.row < h && c.col < w
}
