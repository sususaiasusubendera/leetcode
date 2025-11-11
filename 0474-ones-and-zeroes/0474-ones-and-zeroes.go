func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		cz, co := countOnesZeroes(str)
		for i := m; i >= cz; i-- {
			for j := n; j >= co; j-- {
				dp[i][j] = max(dp[i-cz][j-co]+1, dp[i][j])
			}
		}
	}

	return dp[m][n]
}

// dp (bot-up)
// time: O(lmn)
// space: O(mn)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return (countZeroes, countOnes)
func countOnesZeroes(s string) (int, int) {
	countZeroes, countOnes := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			countZeroes++
			continue
		}
		countOnes++
	}
	return countZeroes, countOnes
}