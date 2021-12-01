package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1() int {
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

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

	err = scanner.Err()
	check(err)

	return increases
}
