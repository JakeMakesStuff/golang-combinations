// Package combinations provides a method to generate all combinations out of a given interface array.
package combinations

import "math/bits"

// All returns all combinations for a given interface array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func All(set []interface{}, f func(subset []interface{})) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []interface{}

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}

		// run function on subset
		f(subset)
	}
}

// Combinations returns combinations of n elements for a given interface array.
// For n < 1, it equals to All and returns all combinations.
func Combinations(set []interface{}, n int, f func(subset []interface{})) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []interface{}

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}

		// run function on subset
		f(subset)
	}
}
