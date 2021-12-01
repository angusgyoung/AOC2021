package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d measurements larger than previous measurement\n", part1())
	fmt.Printf("Part 2: %d measurements larger than previous measurement\n", part2())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isIncrease(depth int, previousDepth int) bool {
	diff := depth - previousDepth

	return diff > 0 && previousDepth != 0
}
