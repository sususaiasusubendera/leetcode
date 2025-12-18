func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	preProfit := make([]int64, n+1) // prefix sum profit
	prePrice := make([]int64, n+1)  // prefix sum price
	for i := 0; i < n; i++ {
		preProfit[i+1] = preProfit[i] + (int64(prices[i]) * int64(strategy[i]))
		prePrice[i+1] = prePrice[i] + int64(prices[i])
	}

	res := preProfit[n]
	for i := k - 1; i < n; i++ {
		leftProfit := preProfit[i-k+1]
		rightProfit := preProfit[n] - preProfit[i+1]
		changeProfit := prePrice[i+1] - prePrice[i-k/2+1]
        res = max(res, leftProfit + changeProfit + rightProfit)
	}

    return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// array, prefix sum, sliding window
// time: O(n)
// space: O(n)