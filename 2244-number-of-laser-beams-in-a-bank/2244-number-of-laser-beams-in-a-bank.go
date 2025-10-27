func numberOfBeams(bank []string) int {
	if len(bank) == 1 {
		return 0
	}

	rows := []int{}
	for i := 0; i < len(bank); i++ {
		totalSecurityDevice := countSecurityDevice(bank[i])
		if totalSecurityDevice > 0 {
			rows = append(rows, totalSecurityDevice)
		}
	}

	ans := 0
	for i := 1; i < len(rows); i++ {
		ans += rows[i] * rows[i-1]
	}

	return ans
}

func countSecurityDevice(rows string) int {
	count := 0
	for i := 0; i < len(rows); i++ {
		if rows[i] == '1' {
			count++
		}
	}

	return count
}

// array, string
// time: O(nm)
// space: O(n)