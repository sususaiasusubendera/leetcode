func minScoreTriangulation(values []int) int {
    n := len(values)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dp func(i, j int) int
    dp = func(i, j int) int { 
        if j-i < 2 { return 0 }

        if memo[i][j] != -1 { return memo[i][j] }

        res := math.MaxInt
        for k := i+1; k < j; k++ {
            res = min(res, dp(i, k) + dp(k, j) + values[i]*values[j]*values[k])
        }
        memo[i][j] = res
        return res
    }

    return dp(0, n-1)
}

// dp (top-down)
// time: O(n^3)
// space: O(n^2)

func min(a, b int) int {
    if a < b { return a }
    return b
}