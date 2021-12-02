package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hpos := 0
	depth := 0

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())

		v, err := strconv.Atoi(command[1])
		check(err)

		switch command[0] {
		case "forward":
			{
				hpos += v
				break
			}
		case "up":
			{
				depth -= v
				break
			}
		case "down":
			{
				depth += v
				break
			}
		}
	}

	return hpos * depth
}
