func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, 102)
	for i := 1; i < 103; i++ {
		dp[i-1] = make([]float64, i)
	}

	dp[0][0] = float64(poured)
	for r := 0; r < query_row+1; r++ {
		for g := 0; g < r+1; g++ {
			e := (dp[r][g] - 1) / 2
			if e > 0 {
				dp[r+1][g] += e
				dp[r+1][g+1] += e
			}
		}
	}

	return min(1, dp[query_row][query_glass])
}

// dp
// time: O(n^2)
// space: O(1)