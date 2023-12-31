package solution_test

import (
	"solution"
	"testing"
)

func TestMakeEqual(t *testing.T) {
	case1 := []string{"abc", "aabc", "bc"}
	if !solution.MakeEqual(case1) {
		t.Error("Expected true, got false")
	}
}
