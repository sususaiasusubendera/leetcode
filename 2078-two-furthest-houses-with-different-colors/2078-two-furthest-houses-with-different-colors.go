func maxDistance(colors []int) int {
	n := len(colors)
	left, right := 0, n-1
	ans := -1
	for left <= right {
		if colors[left] != colors[right] {
			ans = right - left
		}

		// this is so f greedy
		if colors[left] != colors[n-1] || colors[right] != colors[0] {
			ans = right - 0 // "(n - 1) - left" is also valid
		}

        if ans != -1 { return ans }

		left++
		right--
	}

	return 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// array, greedy
// time: O(n)
// space: O(1)