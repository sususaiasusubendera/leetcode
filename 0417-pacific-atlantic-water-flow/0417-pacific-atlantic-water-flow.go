func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // north, east, south, west

	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		stack := [][]int{{r, c}}

		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			visited[curr[0]][curr[1]] = true
			for _, dir := range dirs {
				adj := []int{curr[0] + dir[0], curr[1] + dir[1]}
				newR, newC := adj[0], adj[1]
				if newR >= 0 && newR <= m-1 && newC >= 0 && newC <= n-1 && !visited[newR][newC] && (heights[newR][newC] >= heights[curr[0]][curr[1]]) {
					stack = append(stack, adj)
					visited[newR][newC] = true
				}
			}
		}

	}

	// pacific ocean
	for i := 0; i < m; i++ { dfs(i, 0, pacific) } // left
	for j := 0; j < n; j++ { dfs(0, j, pacific) } // up

	// atlantic ocean
	for i := 0; i < m; i++ { dfs(i, n-1, atlantic) } // right
	for j := 0; j < n; j++ { dfs(m-1, j, atlantic) } // down

	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
            // intersection of (visited) pacific and atlantic
			if pacific[i][j] && atlantic[i][j] { res = append(res, []int{i, j}) }
		}
	}

	return res
}

// dfs, matrix
// time: O(mn)
// space: O(mn)