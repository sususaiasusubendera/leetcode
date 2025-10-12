func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    ans := -1 << 31 // very small number (-1 * 2^(31))
    dp := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        if i >= n - k {
            dp[i] = energy[i]
        } else {
            dp[i] = energy[i] + dp[i + k]
        }

        ans = max(ans, dp[i])
    }

    return ans
}

// dp (bot-up)
// time: O(n)
// space: O(n)

func max(a, b int) int {
    if a > b { return a }

    return b
}