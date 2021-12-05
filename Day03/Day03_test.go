package Day03

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input
var input_data string

//go:embed example
var example_data string

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	example := strings.Split(strings.TrimSuffix(example_data, "\n"), "\n")
	input := strings.Split(strings.TrimSuffix(input_data, "\n"), "\n")

	t.Run("part1", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(198, Part1(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(741950, Part1(input))
		})
	})

	t.Run("part2", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(230, Part2(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(903810, Part2(input))
		})
	})
}
