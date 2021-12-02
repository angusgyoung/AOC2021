package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
