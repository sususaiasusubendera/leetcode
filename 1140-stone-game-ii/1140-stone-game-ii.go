func stoneGameII(piles []int) int {
    n := len(piles)

    // suffix sum
    suffix := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        suffix[i] = piles[i] + suffix[i+1]
    }

    // memo[i][m]
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
    }

    var dp func(i, m int) int
    dp = func(i, m int) int {
        // can take all remaining piles
        if i + (2*m) >= n {
            return suffix[i]
        }

        // already computed
        if memo[i][m] != 0 {
            return memo[i][m]
        }

        bestOpponent := int(^uint(0) >> 1) // max int; minimum stoens for the opponent
        for j := 1; j <= 2*m; j++ {
            opponent := dp(i+j, max(m, j))
            bestOpponent = min(bestOpponent, opponent)
        }

        memo[i][m] = suffix[i] - bestOpponent

        return memo[i][m]
    }

    return dp(0, 1)
}

// array, dp (top-down), prefix sum
// time: O(n^3)
// space: O(n^2)
// notice me senpai, masih ngawang pak
