package Day03

import (
	"aoc2021/helpers"
	"strconv"
)

type accumulator struct {
	// index 0 is the most significant bit
	// positive count means more ones seen
	digit_counts []int
}

func (acc *accumulator) accumulate(digits string) {
	if len(acc.digit_counts) == 0 {
		acc.digit_counts = make([]int, len(digits))
	}

	for i, digit := range digits {
		switch digit {
		case '1':
			acc.digit_counts[i]++
		case '0':
			acc.digit_counts[i]--
		}
	}
}

func (acc *accumulator) gamma() uint {
	var gamma = uint(0)
	for _, digit_count := range acc.digit_counts {
		gamma <<= 1
		if digit_count > 0 {
			gamma |= 1
		}
	}
	return gamma
}

func (acc *accumulator) epsilon() uint {
	return acc.gamma() ^ ((1 << len(acc.digit_counts)) - 1)
}

func Part1(input []string) int {
	acc := accumulator{}
	for _, line := range input {
		acc.accumulate(line)
	}
	return int(acc.epsilon() * acc.gamma())
}

func part2_iterate(input []string, most bool) int {
	possibilities := input
	for i := range input[0] {
		if len(possibilities) <= 1 {
			break
		}

		count := 0
		for _, p := range possibilities {
			switch p[i] {
			case '1':
				count++
			case '0':
				count--
			}
		}

		keep := byte('0')
		if count >= 0 && most || count < 0 && !most {
			keep = byte('1')
		}
		possibilities = helpers.FilterStrings(possibilities,
			func(s string) bool { return s[i] == keep })
	}

	result, _ := strconv.ParseInt(possibilities[0], 2, 32)
	return int(result)
}

func oxygen_generator_rating(input []string) int {
	return part2_iterate(input, true)
}

func co2_scrubber_rating(input []string) int {
	return part2_iterate(input, false)
}

func Part2(input []string) int {
	return oxygen_generator_rating(input) * co2_scrubber_rating(input)
}
