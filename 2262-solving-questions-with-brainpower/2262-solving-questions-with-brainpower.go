func mostPoints(questions [][]int) int64 {
    dp := make([]int64, len(questions)+1)
    for i := len(questions)-1; i >= 0; i-- {
        points, brainpower := questions[i][0], questions[i][1]
        next := i + brainpower + 1

        solve := int64(points)
        if next < len(questions) {
            solve += dp[next]
        }

        skip := dp[i+1]

        if solve > skip {
            dp[i] = solve
        } else {
            dp[i] = skip
        }
    }

    return dp[0]
}

// notice me senpai!!!