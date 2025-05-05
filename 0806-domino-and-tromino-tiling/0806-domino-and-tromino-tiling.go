const MOD = 1000000000+7
func numTilings(n int) int {
    if n == 1 {
        return 1
    }

    dp := [4]int{2, 1, 1, 1}
    for range (n-2) {
        x := (dp[0] + dp[1] + dp[2] + dp[3]) % MOD
        dp[0], dp[1], dp[2], dp[3] = x, dp[2]+dp[3], dp[1]+dp[3], dp[0]
    }

    return dp[0]
}

// NOTICE ME SENPAIII!!!!!!!