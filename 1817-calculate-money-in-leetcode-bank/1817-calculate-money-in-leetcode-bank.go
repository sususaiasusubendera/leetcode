func totalMoney(n int) int {
	weeks := n / 7
	daysRemain := n % 7
	monday := 1
	money := 0
	for monday <= weeks {
		money += (7 * (2*monday + 6)) / 2
		monday++
	}
	money += (daysRemain * (2*monday + daysRemain - 1)) / 2

	return money
}

// math
// time: O(n)
// space: O(1)