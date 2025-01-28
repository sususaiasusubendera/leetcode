var  dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || visited[r][c] {
			return 0
		}
		visited[r][c] = true
		fish := grid[r][c]
		for _, dir := range dirs {
			fish += dfs(r+dir[0], c+dir[1])
		}
		return fish
	}

	maxFish := 0
	for i := range m {
		for j := range n {
			if grid[i][j] > 0 && !visited[i][j] {
				maxFish = max(maxFish, dfs(i, j))
			}
		}
	}

	return maxFish
}

// notice me senpai