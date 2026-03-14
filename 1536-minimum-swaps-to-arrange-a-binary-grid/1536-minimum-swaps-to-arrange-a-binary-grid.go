func minSwaps(grid [][]int) int {
	n := len(grid)
	rightZeros := make([]int, n)
	for i := 0; i < n; i++ {
		countZeros := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				break
			}
			countZeros++
		}
		rightZeros[i] = countZeros
	}

	ops := 0
	for i := 0; i < n; i++ {
		need := n - i - 1

        idx := -1
        for j := i; j < n; j++ {
            if rightZeros[j] >= need {
                idx = j
                break
            }
        }

        if idx == -1 { return idx }

		for idx > i {
			ops++
			rightZeros[idx-1], rightZeros[idx] = rightZeros[idx], rightZeros[idx-1]
			idx--
		}
	}

	return ops
}

// array, greedy, matrix
// time: O(n^2)
// space: O(n^2)