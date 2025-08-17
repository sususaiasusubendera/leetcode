func new21Game(n int, k int, maxPts int) float64 {
    memo := map[int]float64{}
    var dp func(i int) float64
    dp = func(i int) float64 {
        if v, ok := memo[i]; ok {
            return v
        }

        if i >= k {
            if i <= n {
                return 1
            }
            return 0
        }

        prob := 0.0
        for pts := 1; pts <= maxPts; pts++ {
            prob += dp(i + pts)
        }
        res := prob / float64(maxPts)
        memo[i] = res
        return res
    }

    return dp(0)
}

// dp (top-down + memoization)
// TLE
