func solveSudoku(board [][]byte)  {
    n := 9

    row := make([][9]bool, n)
    col := make([][9]bool, n)
    box := make([][9]bool, n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != '.' {
                val := board[i][j] - '1' // 0..8
                bIdx := (i/3)*3 + j/3
                row[i][val], col[j][val], box[bIdx][val] = true, true, true
            }
        }
    }

    var solve func(r, c int) bool
    solve = func(r, c int) bool {
        if r == n {
            return true
        }

        // next cell
        nr, nc := r, c+1
        if nc == n {
            nr++
            nc = 0
        }

        if board[r][c] != '.' {
            return solve(nr, nc)
        }

        bIdx := (r/3)*3 + c/3
        for val := 0; val < 9; val++ {
            if row[r][val] || col[c][val] || box[bIdx][val] {
                continue
            }

            board[r][c] = '1' + byte(val)
            row[r][val], col[c][val], box[bIdx][val] = true, true, true
            if solve(nr, nc) {
                return true
            }

            // backtrack
            board[r][c] = '.'
            row[r][val], col[c][val], box[bIdx][val] = false, false, false
        }

        return false
    }

    solve(0, 0)
}

// array, backtracking, matrix
// time: O(9^81)
// space: O(81)