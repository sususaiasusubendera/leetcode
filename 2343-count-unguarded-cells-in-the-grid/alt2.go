func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    // 0: empty; 1: wall; 2: guard; 3: guarded
    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
    }

    // place the walls
    for _, wall := range walls {
        row, col := wall[0], wall[1]
        grid[row][col] = 1
    }

    // place the guards + check the guarded
    dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left
    for _, guard := range guards {
        row, col := guard[0], guard[1] 
        grid[row][col] = 2 // place a guard
        // check the guard's vision in four directions: north, east, south, and west
        for _, dir := range dirs {
            r, c := row + dir[0], col + dir[1]
            for 0 <= r && r <= m-1 && 0 <= c && c <= n-1 && grid[r][c] != 1 && grid[r][c] != 2 {
                grid[r][c] = 3
                r += dir[0]
                c += dir[1]
            }
        }
    }

    countNotGuarded := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 { countNotGuarded++ }
        }
    }

    return countNotGuarded
}

// array, brute force, matrix, simulation
// time: O(mn + g(m + n))
// space: O(mn)
