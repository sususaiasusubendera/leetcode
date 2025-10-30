func minNumberOperations(target []int) int {
	ans := target[0]
	idx := 1
	for idx < len(target) {
		rising := target[idx] - target[idx-1]
		if rising > 0 {
			ans += rising
		}

		idx++
	}

	return ans
}

// greedy
// time: O(n)
// space: O(1)