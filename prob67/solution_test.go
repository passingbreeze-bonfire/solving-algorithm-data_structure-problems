package main_test

import (
	"github.com/stretchr/testify/assert"
	"prob67/solution"
	"testing"
)

func TestProb66Solution(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(solution.AddBinary("11", "1"), "100", "Case 1 Solved!")
	assert.Equal(solution.AddBinary("1010", "1011"), "10101", "Case 2 Solved!")
	assert.Equal(solution.AddBinary("1", "111"), "1000", "Case 3 Solved!")

}
