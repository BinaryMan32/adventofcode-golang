package Day02

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

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	example := strings.Split(example_data, "\n")
	input := strings.Split(input_data, "\n")

	t.Run("part1", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(150, Part1(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(0, Part1(input))
		})
	})

	t.Run("part2", func(t *testing.T) {
		t.Run("example", func(t *testing.T) {
			assert.Equal(900, Part2(example))
		})
		t.Run("result", func(t *testing.T) {
			assert.Equal(0, Part2(input))
		})
	})
}
