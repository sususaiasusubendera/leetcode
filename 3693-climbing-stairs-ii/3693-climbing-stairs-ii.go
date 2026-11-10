func climbStairs(n int, costs []int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = costs[i-1] + min(dp[i-1]+1, dp[max(0, i-2)]+4, dp[max(0, i-3)]+9)
    }
    return dp[n]
}

// array, dp
// time: O(n)
// space: O(n)