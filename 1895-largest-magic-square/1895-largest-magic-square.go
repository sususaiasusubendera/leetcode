func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// prefix sum horizontal
	preH := make([][]int, m+1)
	for i := range preH {
		preH[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preH[i+1][j+1] = preH[i+1][j] + grid[i][j]
		}
	}

	// prefix sum vertical
	preV := make([][]int, m+1)
	for i := range preV {
		preV[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preV[i+1][j+1] = preV[i][j+1] + grid[i][j]
		}
	}

	ans := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idxI, idxJ := i-1, j-1
			for idxI >= 0 && idxJ >= 0 {
				sum := preH[idxI+1][j+1] - preH[idxI+1][idxJ]
				magic := true

				// horizontal
				for k := idxI; k <= i; k++ {
					currSum := preH[k+1][j+1] - preH[k+1][idxJ]
					if currSum != sum {
						magic = false
						break
					}
				}

				// vertical
				if magic {
					for k := j; k >= idxJ; k-- {
						currSum := preV[i+1][k+1] - preV[idxI][k+1]
						if currSum != sum {
							magic = false
							break
						}
					}
				}

				// diagonal 1
				if magic {
					currSum := 0
					ii, ij := i, j
					for ii >= idxI && ij >= idxJ {
						currSum += grid[ii][ij]
						ii--
						ij--
					}
					if currSum != sum {
						magic = false
					}
				}

				// diagonal 2
				if magic {
					currSum := 0
					ii, ij := i, idxJ
					for ii >= idxI && ij <= j {
						currSum += grid[ii][ij]
						ii--
						ij++
					}
					if currSum != sum {
						magic = false
					}
				}

				if magic {
					side := i - idxI + 1
					ans = max(ans, side)
				}

				idxI--
				idxJ--
			}
		}
	}

	return ans
}

// array, matrix, prefix sum
// time: O(mn * (min(m, n))^2) (gpt)
// space: O(mn)