func maxDotProduct(nums1 []int, nums2 []int) int {
    const NEG_INF = -1 << 63
    n, m := len(nums1), len(nums2)

    // dp[i][j] = best dot product from nums1[i:], nums2[j:]
    dp := make([][]int, n+1)
    for i := range n+1 {
        dp[i] = make([]int, m+1)
        for j := range m+1 {
            dp[i][j] = NEG_INF
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            take := nums1[i] * nums2[j]
            if dp[i+1][j+1] > 0 {
                take += dp[i+1][j+1]
            }

            dp[i][j] = max(take, max(dp[i+1][j], dp[i][j+1])) // take or skip
        }
    }

    return dp[0][0]
}

// array, dp (bot-up)
// time: O(nm)
// space: O(nm)