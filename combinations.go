// Package combinations provides a method to generate all combinations out of a given string array.
package combinations

// All will return all combinations for a given string array
func All(in []string) [][]string {
	length := len(in)
	switch length {
	case 0:
		return [][]string{}
	case 1:
		return [][]string{in}
	}

	maxCount := 1 << uint(length)
	out := [][]string{}
	for i := 1; i < maxCount; i++ {
		item := []string{}
		for j := 0; j < length; j++ {
			if (i & (1 << uint(j))) != 0 {
				item = append(item, in[j])
			}
		}
		out = append(out, item)
	}
	return out
}
