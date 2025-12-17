func maximumProfit(prices []int, k int) int64 {
	memo := make([][][]int64, len(prices))
	for i := range memo {
		memo[i] = make([][]int64, k+1)
		for j := range memo[i] {
			memo[i][j] = make([]int64, 3)
			for s := range memo[i][j] {
				memo[i][j][s] = -1
			}
		}
	}

	// i: day; j: transaction left; s: state
	// state: 
    // - 0: idle
    // - 1: normal transaction (buy - sell)
    // - 2: short selling transaction (sell -> buy)
	var dfs func(int, int, int) int64
	dfs = func(i, j, s int) int64 {
		if j == 0 { // no transaction left
			return 0
		}

		if i == 0 { // day 0
			if s == 0 { // idle
				return 0
			} else if s == 1 { // normal (buy -> sell); OPEN TRANSACTION
				return -int64(prices[0]) // buy
			} else if s == 2 { // short (sell -> buy); OPEN TRANSACTION
				return int64(prices[0]) // sell
			}
		}

		if memo[i][j][s] != -1 {
			return memo[i][j][s]
		}

		p := prices[i]
		res := int64(0)
		if s == 0 { // idle; CLOSE TRANSACTION
			res = max(dfs(i-1, j, 0), max(dfs(i-1, j, 1)+int64(p), dfs(i-1, j, 2)-int64(p)))
		} else if s == 1 { // normal
			res = max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-int64(p))
		} else if s == 2 { // short
			res = max(dfs(i-1, j, 2), dfs(i-1, j-1, 0)+int64(p))
		}

		memo[i][j][s] = res
		return res
	}

	return dfs(len(prices)-1, k, 0)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// array, dp (top-down)
// time: O(nk)
// space: O(nk)