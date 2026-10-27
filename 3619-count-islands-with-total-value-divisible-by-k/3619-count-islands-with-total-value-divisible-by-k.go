func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} // down, right, up, left

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r > m-1 || c < 0 || c > n-1 {
			return 0
		}

		if grid[r][c] == 0 {
			return 0
		}

		total := grid[r][c]
		grid[r][c] = 0 // mark visited
		for _, d := range dirs {
			total += dfs(r+d[0], c+d[1])
		}

		return total
	}

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 && (dfs(i, j)%k == 0) {
                ans++
            }
        }
    }

    return ans
}

// array, dfs, matrix
// time: O(mn)
// space: O(mn)