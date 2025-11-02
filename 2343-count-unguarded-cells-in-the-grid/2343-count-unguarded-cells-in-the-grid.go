func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	// 0: empty; 1: wall; 2: guard; 3: guarded
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	// place the walls
	for _, wall := range walls {
		row, col := wall[0], wall[1]
		grid[row][col] = 1
	}

	// place the guards
	for _, guard := range guards {
		row, col := guard[0], guard[1]
		grid[row][col] = 2
	}

	// horizontal check
	for i := 0; i < m; i++ {
        // left to right
        seen := false
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // wall
				seen = false
			} else if grid[i][j] == 2 { // guard
				seen = true
			} else if seen {
				grid[i][j] = 3 // guarded
			}
		}

        // right to left
        seen = false
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 { // wall
				seen = false
			} else if grid[i][j] == 2 { // guard
				seen = true
			} else if seen {
				grid[i][j] = 3 // guarded
			}
		}
	}

    // vertical check
    for j := 0; j < n; j++ {
        // top to bot
        seen := false
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 { // wall
				seen = false
			} else if grid[i][j] == 2 { // guard
				seen = true
			} else if seen {
				grid[i][j] = 3 // guarded
			}
        }

        // bot to top
        seen = false
        for i := m-1; i >= 0; i-- {
            if grid[i][j] == 1 { // wall
				seen = false
			} else if grid[i][j] == 2 { // guard
				seen = true
			} else if seen {
				grid[i][j] = 3 // guarded
			}
        }
    }

	countNotGuarded := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				countNotGuarded++
			}
		}
	}

	return countNotGuarded
}

// array, matrix, simulation
// time: O(mn)
// space: O(mn)