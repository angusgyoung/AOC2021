package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)

	previousDepth := 0
	increases := 0

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)

		if isIncrease(depth, previousDepth) {
			increases++
		}

		previousDepth = depth
	}

	err := scanner.Err()
	check(err)

	return increases
}

// Return if increase in depth and previous depth not 0
func isIncrease(depth int, previousDepth int) bool {
	diff := depth - previousDepth

	return diff > 0 && previousDepth != 0
}
