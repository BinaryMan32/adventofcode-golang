package helpers

import (
	"strconv"
	"strings"
)

func ReadNumbers(data string) []int {
	lines := strings.Split(data, "\n")
	nums := make([]int, 0, len(lines))

	for _, x := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(x) == 0 {
			continue
		}
		n, _ := strconv.Atoi(x)
		nums = append(nums, n)
	}

	return nums
}

func FilterStrings(input []string, keep func(string) bool) []string {
	filtered := make([]string, 0)
	for _, item := range input {
		if keep(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
