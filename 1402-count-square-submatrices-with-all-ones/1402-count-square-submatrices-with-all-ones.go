func countSquares(matrix [][]int) int {
    rows, cols := len(matrix), len(matrix[0])
    dp := make([][]int, rows)
    for i := range dp {
        dp[i] = make([]int, cols)
    }

    res := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == 1 {
                if i == 0 || j == 0 {
                    dp[i][j] = 1 // basecase for the first row/col
                } else {
                    // check dp left, left-up (diagonal), and up
                    dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
                }
                res += dp[i][j]
            }
        }
    }

    return res
}

// dp (bottom-up)
// time: O(mn)
// space: O(mn)

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        }
        return c
    }

    if b < c {
        return b
    }
    return c
}
