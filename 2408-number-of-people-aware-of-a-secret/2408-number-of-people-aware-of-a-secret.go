func peopleAwareOfSecret(n int, delay int, forget int) int {
    const MOD = 1e9 + 7
    dp := make([]int, n+1)
    dp[1] = 1
    for i := 1; i <= n; i++ {
        for j := i + delay; j <= i + forget - 1 && j <= n; j++ {
            dp[j] = (dp[j] + dp[i]) % MOD
        }
    }

    ans := 0
    for i := n - forget + 1; i <= n; i++ {
        if i >= 1 {
            ans = (ans + dp[i]) % MOD
        }
    }

    return ans
}

// dp (bot-up)
// time: O(n)
// space: O(n)