package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part2() int {
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hpos := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())

		v, err := strconv.Atoi(command[1])
		check(err)

		switch command[0] {
		case "forward":
			{
				hpos += v
				depth += (aim * v)
				break
			}
		case "up":
			{
				aim -= v
				break
			}
		case "down":
			{
				aim += v
				break
			}
		}
	}

	return hpos * depth
}
