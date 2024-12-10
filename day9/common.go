package day9

import "strconv"

// expand compacted format to get representation of file blocks and free blocks
func expandBlocks(data string) []string {
	expanded := []string{}
	for i := range data {
		blocks, err := strconv.Atoi(string(data[i]))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			fileId := i / 2
			for j := 0; j < blocks; j++ {
				expanded = append(expanded, strconv.Itoa(fileId))
			}
		} else {
			for j := 0; j < blocks; j++ {
				expanded = append(expanded, ".")
			}
		}
	}

	return expanded
}

func computeChecksum(blocks []string) int {
	sum := 0
	for i, block := range blocks {
		if block == "." {
			continue
		}

		val, err := strconv.Atoi(block)
		if err != nil {
			panic(err)
		}

		sum += i * val
	}

	return sum
}
