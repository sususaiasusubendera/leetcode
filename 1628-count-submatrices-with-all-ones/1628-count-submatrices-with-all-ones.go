func numSubmat(mat [][]int) int {
    res := 0
    rows, cols := len(mat), len(mat[0])
    height := make([]int, len(mat[0]))
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if mat[i][j] == 1 {
                height[j]++
            } else {
                height[j] = 0
                continue
            }

            minH := height[j]
            for k := j; k >= 0 && height[k] > 0; k-- {
                if height[k] < minH {
                    minH = height[k]
                }
                res += minH
            }
        }
    }

    return res
}

// dp
// time: O(rows * cols^2)
// space: O(rows * cols)

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}