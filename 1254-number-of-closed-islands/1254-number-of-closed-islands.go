func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// mark all the open islands
	// top
	for c := 0; c < n; c++ {
		if grid[0][c] == 0 {
			bfs(0, c, grid)
		}
	}

	// right
	for r := 0; r < m; r++ {
		if grid[r][n-1] == 0 {
			bfs(r, n-1, grid)
		}
	}

	// bot
	for c := n - 1; c >= 0; c-- {
		if grid[m-1][c] == 0 {
			bfs(m-1, c, grid)
		}
	}

	// left
	for r := m - 1; r >= 0; r-- {
		if grid[r][0] == 0 {
			bfs(r, 0, grid)
		}
	}

	// count all the closed islands
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				bfs(i, j, grid)
				count++
			}
		}
	}

	return count
}

func bfs(x, y int, grid [][]int) {
	m, n := len(grid), len(grid[0])
	grid[x][y] = 1 // mark
	q := [][]int{{x, y}}
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up, right, down, left
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, d := range dirs {
			nx, ny := curr[0]+d[0], curr[1]+d[1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == 0 {
                grid[nx][ny] = 1
				q = append(q, []int{nx, ny})
			}
		}
	}
}

// array, bfs
// time: O(mn)
// space: O(mn)