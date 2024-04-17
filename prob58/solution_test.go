package main_test

import (
	"github.com/stretchr/testify/assert"
	"prob58/solution"
	"testing"
)

func TestProb58Solution(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(solution.LengthOfLastWord("Hello World"), 5, "Case 1 Solved!")
	assert.Equal(solution.LengthOfLastWord("   fly me   to   the moon  "), 4, "Case 2 Solved!")
	assert.Equal(solution.LengthOfLastWord("luffy is still joyboy"), 6, "Case 3 Solved!")
}
