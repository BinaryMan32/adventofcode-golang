package Day01

func Part1(input []int) int {
	count := 0
	for i, next := range input[1:] {
		prev := input[i]
		if next > prev {
			count++
		}
	}
	return count
}

func Part2(input []int) int {
	count := 0
	for i, next := range input[3:] {
		prev := input[i]
		if next > prev {
			count++
		}
	}
	return count
}
