package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

	fmt.Printf("Part 1: %d measurements larger than previous measurement\n", part1(file))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
