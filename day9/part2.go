package day9

import (
	"bufio"
	"fmt"
	"os"
)

// similar to part one, but only move full files at a time so no fragmentation is
// introduced. Means we need to check if there is space for the whole file
func compactFilesNoFragmentation(blocks []string) []string {
	compacted := make([]string, len(blocks))
	copy(compacted, blocks)

	// keep trying until we went through all the blocks and were unable to move any files
	currentFileIndices := []int{}

	for i := len(compacted) - 1; i >= 0; i-- {
		if compacted[i] == "." {
			continue
		}

		if len(currentFileIndices) == 0 || compacted[i] == compacted[currentFileIndices[0]] {
			currentFileIndices = append(currentFileIndices, i)
		} else {
			emptyBlockIndices := []int{}
			for j, b := range compacted {
				if b != "." {
					if j >= currentFileIndices[0] || len(emptyBlockIndices) >= len(currentFileIndices) {
						// if we went past current file indices, stop searching because we can't shift files to the right
						// continue Loop
						// if we hit a non empty space and slice length is greater than or equal to file block size,
						// end loop - we've found a space to move the file blocks too
						break
					} else {
						// otherwise, reset empty block indices, we'll start trying to find a new empty block
						emptyBlockIndices = []int{}
						continue
					}
				}

				emptyBlockIndices = append(emptyBlockIndices, j)
			}

			// move file if it fits
			if len(emptyBlockIndices) >= len(currentFileIndices) {
				for j := range currentFileIndices {
					idxToReplace := emptyBlockIndices[j]
					idxOfReplacement := currentFileIndices[j]
					compacted[idxToReplace] = compacted[idxOfReplacement]
					compacted[idxOfReplacement] = "."
				}
			}

			currentFileIndices = []int{i}
		}
	}

	return compacted
}

func SolvePart2() {
	file, err := os.Open("input/day9input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Text()

	expanded := expandBlocks(data)
	compacted := compactFilesNoFragmentation(expanded)
	checksum := computeChecksum(compacted)

	fmt.Println(checksum)
}
