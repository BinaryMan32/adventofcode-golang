package Day01

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readNumbers(fileName string) []int {
	content, _ := ioutil.ReadFile("testdata/" + fileName)
	lines := strings.Split(string(content), "\n")
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

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	example := readNumbers("example")
	input := readNumbers("input")

	t.Run("part1", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(7, Part1(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(0, Part1(input))
		})
	})

	t.Run("part2", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(5, Part2(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(0, Part2(input))
		})
	})
}
