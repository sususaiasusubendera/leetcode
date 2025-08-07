func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }

    maxLen := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }

        maxLen = max(maxLen, dp[i])
    }

    return maxLen
}

// dp
// notice me senpai???

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}