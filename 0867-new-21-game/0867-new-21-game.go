func new21Game(n int, k int, maxPts int) float64 {
    if k == 0 || n >= k+maxPts {
        return 1.0
    }

    dp := make([]float64, n+1)
    dp[0] = 1.0
    windowSum := 1.0
    res := 0.0
    for i := 1; i <= n; i++ {
        dp[i] = windowSum / float64(maxPts)
        
        if i < k {
            windowSum += dp[i]
        } else {
            res += dp[i]
        }

        if i-maxPts >= 0 {
            windowSum -= dp[i-maxPts]
        }
    }

    return res
}

// dp (bottom-up), sliding window
// time: O(n)
// space: O(n)