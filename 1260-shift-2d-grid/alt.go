func shiftGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])
  n := r * c
	k = k % n

    if k == 0 { return grid }

    rev(grid, 0, n-1)
    rev(grid, 0, k-1)
    rev(grid, k, n-1)

    return grid
}

func rev(grid [][]int, i, j int) {
    c := len(grid[0])
    for i < j {
        grid[i/c][i%c], grid[j/c][j%c] = grid[j/c][j%c], grid[i/c][i%c]
        i++
        j--
    }
}

// array, simulation
// time: O(n)
// space: O(1)

// this approach is (prolly) the most optimal one
