func bestClosingTime(customers string) int {
	// close at hour 0
	penaltyy := 0 // penalty y
	for _, c := range customers {
		if c == 'Y' {
			penaltyy++
		}
	}

	totalPenalty := penaltyy
	ans := 0 // idx hour with min penalty
	y := 0
	penaltyn := 0 // penalty n
	for i, c := range customers {
		if c == 'Y' {
			y++
		} else {
			penaltyn++
		}

		currPenalty := penaltyy - y + penaltyn
		if currPenalty < totalPenalty {
			totalPenalty = currPenalty
			ans = i + 1
		}
	}

	return ans
}

// prefix sum, string
// time: O(n)
// space: O(1)