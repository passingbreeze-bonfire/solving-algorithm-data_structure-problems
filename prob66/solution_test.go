package main_test

import (
	"github.com/stretchr/testify/assert"
	"prob66/solution"
	"testing"
)

func TestProb66Solution(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(solution.PlusOne([]int{1, 2, 3}), []int{1, 2, 4}, "Case 1 Solved!")
	assert.Equal(solution.PlusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2}, "Case 2 Solved!")
	assert.Equal(solution.PlusOne([]int{9}), []int{1, 0}, "Case 3 Solved!")
}
