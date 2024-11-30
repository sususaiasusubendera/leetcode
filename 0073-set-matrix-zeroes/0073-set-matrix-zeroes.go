func setZeroes(matrix [][]int)  {
    // check 0 for row 0 and col 0
    rz, cz := false, false
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            cz = true
            break
        }
    }
    for j := 0; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            rz = true
            break
        }
    }

    // check 0 in matrix, except for row 0 and col 0, mark if any
    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[0][j], matrix[i][0] = 0, 0
            }
        }
    }

    // check mark 0 in row 0 and col 0
    for i := 1; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < len(matrix[0]); j++ {
                matrix[i][j] = 0
            }
        }
    }
    for j := 1; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            for i := 0; i < len(matrix); i++ {
                matrix[i][j] = 0
            }
        }
    }

    // check the initial state for row 0 and col 0
    if rz {
        for j := 0; j < len(matrix[0]); j++ {
            matrix[0][j] = 0
        }
    }
    if cz {
        for i := 0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }
}

// rz: check if there is 0 in row 0
// cz: check if there is 0 in col 0