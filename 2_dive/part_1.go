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

	x := 0
	y := 0

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())

		v, err := strconv.Atoi(command[1])
		check(err)

		switch command[0] {
		case "forward":
			{
				x += v
				break
			}
		case "up":
			{
				y -= v
				break
			}
		case "down":
			{
				y += v
				break
			}
		}
	}

	return x * y
}
