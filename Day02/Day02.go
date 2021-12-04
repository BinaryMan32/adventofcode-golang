package Day02

import "fmt"

func Part1(input []string) int {
	var (
		x       int = 0
		y       int = 0
		command string
		arg     int
	)
	for _, line := range input {
		n, _ := fmt.Sscanf(line, "%s %d", &command, &arg)
		if n < 2 {
			continue
		}
		switch command {
		case "forward":
			x += arg
		case "down":
			y += arg
		case "up":
			y -= arg
		}
	}
	return x * y
}

func Part2(input []string) int {
	var (
		x       int = 0
		y       int = 0
		aim     int = 0
		command string
		arg     int
	)
	for _, line := range input {
		n, _ := fmt.Sscanf(line, "%s %d", &command, &arg)
		if n < 2 {
			continue
		}
		switch command {
		case "forward":
			x += arg
			y += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		}
	}
	return x * y
}
