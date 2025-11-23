func maxSumDivThree(nums []int) int {
	dp := []int{0, -1, -1} // remainder x % 3 (0, 1, 2); sum = 0 -> r = 0
	for _, num := range nums {
		prev := make([]int, 3)
		copy(prev, dp)
		for r := 0; r < 3; r++ {
			newR := (r + num) % 3
			if prev[r] != -1 { //only valid prev[r] (total for "remainder = r" state)
				dp[newR] = max(dp[newR], prev[r]+num)
			}
		}
	}

	return dp[0]
}

// dp (bot-up)
// time: O(n)
// space: O(1)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}