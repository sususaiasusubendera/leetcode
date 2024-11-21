func isValidSudoku(board [][]byte) bool {
    r := make([]map[byte]bool, 9)
    c := make([]map[byte]bool, 9)
    b := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        r[i] = map[byte]bool{}
        c[i] = map[byte]bool{}
        b[i] = map[byte]bool{}
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            n := board[i][j]
            if n == '.' {
                continue
            }

            bIndex := (i/3)*3 + j/3
            if r[i][n] || c[j][n] || b[bIndex][n] {
                return false
            }

            r[i][n] = true
            c[j][n] = true
            b[bIndex][n] = true
        }
    }
    return true
}