func minimumArea(grid [][]int) int {
    top, bot := 1000, 0
    left, right := 1000, 0
    rows, cols := len(grid), len(grid[0])
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                if i < top {
                    top = i
                }
                if i > bot {
                    bot = i
                }
                if j < left {
                    left = j
                }
                if j > right {
                    right = j
                }
            }
        }
    }

    return (bot-top+1) * (right-left+1)
}

// array, matrix
// time: O(mn)
// space: O(1)