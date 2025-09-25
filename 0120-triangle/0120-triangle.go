func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, len(triangle[n-1]))

    for i := 0; i < len(triangle[n-1]); i++ {
        dp[i] = triangle[n-1][i]
    }

    for i := n-2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
        }
    }

    return dp[0]
}

// dp (bot-up)
// time: O(n^2)
// space: O(n)

func min(a, b int) int {
    if a < b { return a }
    return b
}