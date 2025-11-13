func maxOperations(s string) int {
	ops := 0
	idx := 0
	countOnes := 0
	for idx < len(s) {
		if idx+1 < len(s) && s[idx] == '1' && s[idx+1] == '0' {
			countOnes++
			for idx+1 < len(s) && s[idx+1] == '0' {
				idx++
			}
			ops += countOnes
		} else if s[idx] == '1' {
			countOnes++
			idx++
		} else {
            // s[idx] == '0'
			idx++
		}
	}

	return ops
}

// counting, greedy, string
// time: O(n)
// space: O(1)