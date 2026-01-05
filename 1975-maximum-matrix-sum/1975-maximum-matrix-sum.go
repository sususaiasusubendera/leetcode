func maxMatrixSum(matrix [][]int) int64 {
    sum := 0
    countNeg := 0
    minPos, minNeg := 100_000, 100_000
    zeroExists := false
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            sum += abs(matrix[i][j])

            if matrix[i][j] < 0 {
                countNeg++
                minNeg = min(minNeg, -matrix[i][j])
            } else if matrix[i][j] == 0 {
                zeroExists = true
            } else if matrix[i][j] > 0 {
                minPos = min(minPos, matrix[i][j])
            }
        }
    }

    if countNeg%2 == 0 || zeroExists {
        return int64(sum)
    }

    return int64(sum - 2*min(minPos, minNeg))
}

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

// array, greedy, matrix
// time: O(n^2)
// space: O(1)