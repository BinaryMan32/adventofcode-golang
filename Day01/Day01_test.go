package Day01

import (
	_ "embed"
	"testing"

	"aoc2021/helpers"

	"github.com/stretchr/testify/assert"
)

//go:embed input
var input_data string

//go:embed example
var example_data string

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	example := helpers.ReadNumbers(example_data)
	input := helpers.ReadNumbers(input_data)

	t.Run("part1", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(7, Part1(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(1521, Part1(input))
		})
	})

	t.Run("part2", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(5, Part2(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(1543, Part2(input))
		})
	})
}
