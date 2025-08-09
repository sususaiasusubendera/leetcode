func soupServings(n int) float64 {
    serving := (n + 24) / 25 // ceiling
    memo := map[[2]int]float64{}

    var dp func(i, j int) float64
    dp = func(i, j int) float64 {
        if i <= 0 && j <= 0 {
            return 0.5
        }
        if i <= 0 {
            return 1
        }
        if j <= 0 {
            return 0
        }

        key := [2]int{i, j}
        if val, ok := memo[key]; ok {
            return val
        }

        memo[key] = (dp(i-4, j) + dp(i-3, j-1) + dp(i-2, j-2) + dp(i-1, j-3)) / 4
        return memo[key]
    }

    // optimization
    for i := 1; i <= serving; i++ {
        if dp(i, i) > 1 - 1e-5 {
            return 1
        }
    }

    return dp(serving, serving)
}

// dp
// time: O(n^2)
// space: O(n^2)