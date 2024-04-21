package solution_test

import (
	"solving-algorithm-data_structure-problems/prob2610/solution"
	"testing"
)

func TestMakeEqual(t *testing.T) {
	case1 := []string{"abc", "aabc", "bc"}
	if !solution.MakeEqual(case1) {
		t.Error("Case1 Expected true, got false")
	}

	case2 := []string{"ab", "a"}
	if solution.MakeEqual(case2) {
		t.Error("Case2 Expected false, got true")
	}
}
