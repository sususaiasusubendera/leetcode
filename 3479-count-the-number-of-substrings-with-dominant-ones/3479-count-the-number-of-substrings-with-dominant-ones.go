func numberOfSubstrings(s string) int {
	// build a prefix sum pre
	// pre[i] stores the index of the nearest '0' at or before position (i-1)
	pre := make([]int, len(s)+1)
	pre[0] = -1
	for i := 0; i < len(s); i++ {
		if i == 0 || (i > 0 && s[i-1] == '0') {
			pre[i+1] = i
		} else {
			pre[i+1] = pre[i]
		}
	}

	res := 0
	// i: "right" index of substring (1-indexed)
	for i := 1; i <= len(s); i++ {
		count0 := 0
		if s[i-1] == '0' { count0++ }

		// j: "left" index of substring (1-indexed)
		j := i
		for j > 0 && count0*count0 <= len(s) {
			// count1 = window (pre[j], i) - count0
			count1 := (i - pre[j]) - count0
			if count0*count0 <= count1 {
				add := j - pre[j] // all possible starting position (pre[j]) from (pre[j], j)
				// limit if the window is too large for minimal valid length
				add = min(add, count1-count0*count0+1)
				res += add
			}
			j = pre[j]
			count0++
		}
	}
	return res
}

// prefix sum, sliding window, string
// time: O(n^2)
// space: O(n)

func min(a, b int) int {
	if a < b { return a }
	return b
}