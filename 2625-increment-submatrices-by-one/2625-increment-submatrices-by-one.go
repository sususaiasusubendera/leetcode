func rangeAddQueries(n int, queries [][]int) [][]int {
    diff := make([][]int, n+1)
    for i := range diff {
        diff[i] = make([]int, n+1)
    }
    prefix := make([][]int, n+1)
    for i := range prefix {
        prefix[i] = make([]int, n+1)
    }
    
    // build difference matrix
    for _, query := range queries {
        r1, c1, r2, c2 := query[0], query[1], query[2], query[3]
        diff[r1][c1] += 1
        diff[r1][c2+1] -= 1
        diff[r2+1][c1] -= 1
        diff[r2+1][c2+1] += 1
    }

    // build prefix sum matrix
    for i := 0; i < n+1; i++ {
        for j := 0; j < n+1; j++ {
            if i == 0 && j == 0 {
                prefix[i][j] += diff[i][j]
            } else if i == 0 {
                prefix[i][j] += diff[i][j] + prefix[i][j-1]
            } else if j == 0 {
                prefix[i][j] += diff[i][j] + prefix[i-1][j]
            } else {
                prefix[i][j] += diff[i][j] + prefix[i][j-1] + prefix[i-1][j] - prefix[i-1][j-1]
            }
        }
    }

    // copy prefix to ans (size n x n)
    ans := make([][]int, n)
    for i := 0; i < n; i++ {
        ans[i] = make([]int, n)
        copy(ans[i], prefix[i][:n+1])
    }

    return ans
}

// array, difference array 2D, matrix, prefix sum
// time: O(n^2)
// space: O(n^2)