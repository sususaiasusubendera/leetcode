func shiftGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])
	n := r * c
	k = k % n

	if k == 0 {
		return grid
	}

	temp := make([]int, n)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp[(i*c)+j] = grid[i][j]
		}
	}

    temp = append(temp[len(temp)-k:], temp[:len(temp)-k]...)

    for i := 0; i < len(temp); i++ {
        grid[i/c][i%c] = temp[i]
    }

    return grid
}

// array
// time: O(n)
// space: O(n)