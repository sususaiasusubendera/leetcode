func numSub(s string) int {
	idx := 0
	res := 0
	for idx < len(s) {
		if s[idx] == '1' {
            countOne := 0
			for idx < len(s) && s[idx] == '1' {
				countOne++
				idx++
			}
            res += countSubstring(countOne)
            continue
		}
        idx++
	}

    return res % (1e9 + 7)
}

// math, string
// time: O(n)
// space: O(1)

func countSubstring(n int) int {
	return ((n + 1) * n) / 2
}