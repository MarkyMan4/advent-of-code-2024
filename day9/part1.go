package day9

import (
	"bufio"
	"fmt"
	"os"
)

// move blocks from the rightmost to the leftmost free block, until nothing more can be moved
func compactFilesFragmented(blocks []string) []string {
	compacted := make([]string, len(blocks))
	copy(compacted, blocks)

Loop:
	for i, block := range compacted {
		if block == "." {
			for j := len(compacted) - 1; j >= 0; j-- {
				if j <= i {
					// if j and i meet, nothing more can be moved, break outer loop
					break Loop
				}

				// swap something from the right with empty space on left, and break inner loop, continue search from left
				if compacted[j] != "." {
					compacted[i] = compacted[j]
					compacted[j] = "."
					break
				}
			}
		}
	}

	return compacted
}

func SolvePart1() {
	file, err := os.Open("input/day9input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Text()

	expanded := expandBlocks(data)
	compacted := compactFilesFragmented(expanded)
	checksum := computeChecksum(compacted)

	fmt.Println(checksum)
}
