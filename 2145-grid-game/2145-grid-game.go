func gridGame(grid [][]int) int64 {
    n := len(grid[0])

    // precompute prefix sums for row0 and row1
    preRow0, preRow1 := make([]int, n), make([]int, n)
    preRow0[0], preRow1[0] = grid[0][0], grid[1][0]
    for i := 1; i < n; i++ {
        preRow0[i] = preRow0[i-1] + grid[0][i]
        preRow1[i] = preRow1[i-1] + grid[1][i]
    }

    // minimize the maximum points collected by robot 2
    minMaxPoints := int64(1<<63 - 1) // large initial value
    for i := 0; i < n; i++ {
        // points collected by robot 2 if robot 1 moves down at col i
        top := int64(preRow0[n-1] - preRow0[i])
        var bot int64
        if i == 0 {
            bot = int64(0) // no points i row 1 before col 0
        } else {
            bot = int64(preRow1[i-1])
        }

        // robot 2 will take the maximum points from top or bot
        maxPoints := max(top, bot)

        // minimize the maximum points robot 2 can take
        if maxPoints < minMaxPoints {
            minMaxPoints = maxPoints
        }
    }

    return minMaxPoints
}

// time: O(n)
// space: O(n)

func max(a, b int64) int64 {
    if a > b {
        return a
    } else {
        return b
    }
}