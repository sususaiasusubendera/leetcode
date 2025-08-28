func sortMatrix(grid [][]int) [][]int {
    rows, cols := len(grid), len(grid[0])
    for i := rows-1; i >= 0; i-- {
        if i == rows-1 { continue }

        x, y := i, 0
        temp := []int{}
        for x <= rows-1 {
            temp = append(temp, grid[x][y])
            x++
            y++
        }

        slices.SortFunc(temp, func(a, b int) int { return b - a })
        x, y = i, 0
        for x <= rows-1 {
            grid[x][y] = temp[y]
            x++
            y++
        }
    }

    for j := 1; j < cols; j++ {
        if j == cols-1 { continue }

        x, y := 0, j
        temp := []int{}
        for y <= cols-1 {
            temp = append(temp, grid[x][y])
            x++
            y++
        }

        slices.Sort(temp)
        x, y = 0, j
        for y <= cols-1 {
            grid[x][y] = temp[x]
            x++
            y++
        }
    }

    return grid
}

// array, matrix
// time: O(mn.log(min(mn)))
// space: O(min(mn))