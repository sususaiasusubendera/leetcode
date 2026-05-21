func longestCommonPrefix(arr1 []int, arr2 []int) int {
	ans := 0
	m1 := map[int]bool{}

	for _, a1 := range arr1 {
		for a1 > 0 && !m1[a1] {
			m1[a1] = true
			a1 /= 10
		}
	}

	for _, a2 := range arr2 {
		for a2 > 0 && !m1[a2] {
			a2 /= 10
		}

		if a2 == 0 {
			continue
		}

		ans = max(ans, len(strconv.Itoa(a2)))
	}

	return ans
}

// array, hash map

// #1 2026/05/21