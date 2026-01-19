func maxSideLength(mat [][]int, threshold int) int {
    m, n := len(mat), len(mat[0])

    // prefix sum 2d
    pre := make([][]int, m+1)
    for i := range pre {
        pre[i] = make([]int, n+1)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + mat[i][j]
        }
    }

    ans := 0
    sub := 0 // substractor -> top left
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 {
                if mat[i][j] <= threshold {
                    ans = max(ans, 1)
                }
                continue
            }

            idxI, idxJ := i-sub-1, j-sub-1
            tempSub := sub
            for idxI >= 0 && idxJ >= 0 {
                tempSub++
                sum := pre[i+1][j+1] - pre[idxI][j+1] - pre[i+1][idxJ] + pre[idxI][idxJ]
                if sum <= threshold {
                    side := i - idxI + 1
                    ans = max(ans, side)
                    sub = max(sub, tempSub)
                }
                idxI--
                idxJ--
            }
        }
    }

    return ans
}

// array, matrix, prefix sum
// time: O(mn * min(m, n))
// space: O(mn)