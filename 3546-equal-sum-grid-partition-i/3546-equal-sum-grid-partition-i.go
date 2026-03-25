func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    pre := make([][]int, m+1)
    for i := range pre {
        pre[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            pre[i][j] = grid[i-1][j-1] + pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] 
        }
    }

    // horizontal line
    if m > 1 {
        for i := 0; i < m-1; i++ {
            up, down := pre[i+1][n], pre[m][n] - pre[i+1][n]
            if up == down { return true }
        }
    }

    // vertical
    if n > 1 {
        for j := 0; j < n-1; j++ {
            left, right := pre[m][j+1], pre[m][n] - pre[m][j+1]
            if left == right { return true }
        }
    }

    return false
}

// array, prefix sum, matrix
// time: O(mn)
// space: O(mn)