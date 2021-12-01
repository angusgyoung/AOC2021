package main

import (
	"bufio"
	"os"
	"strconv"
)

func part2() int {
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	previousDepthSum := 0
	increases := 0

	var window []int

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)

		window = append(window, depth)

		if len(window) == 3 {
			depthSum := sum(window...)

			if isIncrease(depthSum, previousDepthSum) {
				increases++
			}

			_, window = window[0], window[1:]
			previousDepthSum = depthSum
		}
	}

	return increases
}

func sum(depths ...int) int {
	result := 0
	for _, i := range depths {
		result += i
	}
	return result
}
