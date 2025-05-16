func getWordsInLongestSubsequence(words []string, groups []int) []string {
    n := len(groups)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxIndex := 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if check(words[i], words[j]) && dp[j]+1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	ans := []string{}
	for i := maxIndex; i >= 0; i = prev[i] {
		ans = append(ans, words[i])
	}
	reverse(ans)
	return ans
}

// editorial
// DP
// NOTICE ME SENPAI!!!

func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
            if diff > 1 {
                return false
            }
		}
	}
	return diff == 1
}

func reverse(arr []string) {
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}