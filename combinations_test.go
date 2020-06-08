package combinations

import "testing"

func TestCombinations(t *testing.T) {
	subsets := make([][]interface{}, 0, 2)
	test := []interface{}{uint8(0), uint8(1), uint8(2), uint8(3)}
	Combinations(test, 2, func(subset []interface{}) bool {
		subsets = append(subsets, subset)
		return false
	})
	if len(subsets) != 6 {
		t.Error("subsets returned was len of", len(subsets))
		t.FailNow()
		return
	}
}

func TestAll(t *testing.T) {
	subsets := make([][]interface{}, 0, 2)
	test := []interface{}{uint8(0), uint8(1), uint8(2), uint8(3)}
	All(test, func(subset []interface{}) bool {
		subsets = append(subsets, subset)
		return false
	})
	if len(subsets) != 15 {
		t.Error("subsets returned was len of", len(subsets))
		t.FailNow()
		return
	}
}
