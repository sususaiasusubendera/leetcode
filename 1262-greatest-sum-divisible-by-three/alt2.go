func maxSumDivThree(nums []int) int {
	dp := make([][]int, len(nums)+1)
	dp[0] = []int{0, -1, -1} // remainder x % 3 (0, 1, 2); sum = 0 -> r = 0

	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		rem := num % 3
		dp[i] = []int{-1, -1, -1}
		for r := 0; r < 3; r++ {
			// skip num
			if dp[i-1][r] != -1 {
				dp[i][r] = max(dp[i][r], dp[i-1][r])
			}

			// take num
			if dp[i-1][r] != -1 {
				newR := (r + rem) % 3
				dp[i][newR] = max(dp[i][newR], dp[i-1][r]+num)
			}

		}
	}

	return dp[len(nums)][0]
}

// dp (bot-up)
// time: O(n)
// space: O(n)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
