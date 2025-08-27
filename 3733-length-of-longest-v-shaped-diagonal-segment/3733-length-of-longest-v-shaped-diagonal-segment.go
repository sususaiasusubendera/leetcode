func lenOfVDiagonal(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    longest := 0
    // dirs: up-right, down-right, down-left, up-left
    dirs := [][2]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
    turnDirs := map[int]int{
        0: 1,
        1: 2,
        2: 3,
        3: 0,
    }
    memo := make([][][][]int, rows)
    for i := 0; i < rows; i++ {
        memo[i] = make([][][]int, cols)
        for j := 0; j < cols; j++ {
            memo[i][j] = make([][]int, 4) // dirs
            for d := 0; d < 4; d++ {
                memo[i][j][d] = []int{-1, -1} // isTurned: 1 true, 0 false
            }
        }
    }

    var dp func(x, y, dirIdx, isTurned, target int) int
    dp = func(x, y, dirIdx, isTurned, target int) int {
        if x < 0 || y < 0 || x > rows-1 || y > cols-1 {
            return 0
        }
        if grid[x][y] != target {
            return 0
        }

        if memo[x][y][dirIdx][isTurned] != -1 {
            return memo[x][y][dirIdx][isTurned]
        }

        nx, ny := x+dirs[dirIdx][0], y+dirs[dirIdx][1]
        nextTarget := changeTarget(target)

        res := dp(nx, ny, dirIdx, isTurned, nextTarget)
        if isTurned == 0 {
            res = max(res, dp(nx, ny, turnDirs[dirIdx], 1, nextTarget))
        }

        memo[x][y][dirIdx][isTurned] = 1 + res
        return memo[x][y][dirIdx][isTurned]
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                for dirIdx := range dirs {
                    longest = max(longest, dp(i, j, dirIdx, 0, 1))
                }
            }
        }
    }

    return longest
}

// dp (top-down + memoization)
// time: O(mn)
// space: O(mn)

func changeTarget(target int) int {
    if target == 1 || target == 0 {
        return 2
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}