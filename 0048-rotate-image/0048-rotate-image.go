func rotate(matrix [][]int)  {
    rowLen, colLen := len(matrix), len(matrix[0])
    for i := 0; i < rowLen; i++ {
        // transpose the matrix
        for j := i; j < colLen; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }

        // flip the row
        for left, right := 0, colLen - 1; left < right; left, right = left + 1, right - 1 {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
        }
    }
}