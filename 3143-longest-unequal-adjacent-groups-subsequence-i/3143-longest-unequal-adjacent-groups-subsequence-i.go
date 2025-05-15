func getLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	maxLen, endIndex := 1, 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}
	for i := 1; i < n; i++ {
		bestLen, bestPrev := 1, -1
		for j := i - 1; j >= 0; j-- {
			if groups[i] != groups[j] && dp[j]+1 > bestLen {
				bestLen, bestPrev = dp[j]+1, j
			}
		}
		dp[i] = bestLen
		prev[i] = bestPrev
		if dp[i] > maxLen {
			maxLen, endIndex = dp[i], i
		}
	}

	res := make([]string, 0)
	for i := endIndex; i != -1; i = prev[i] {
		res = append(res, words[i])
	}
	reverse(res)
	return res
}

// notice me senpai

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}