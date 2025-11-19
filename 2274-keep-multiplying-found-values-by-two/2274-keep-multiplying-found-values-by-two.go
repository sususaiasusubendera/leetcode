func findFinalValue(nums []int, original int) int {
	m := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, exists := m[nums[i]]; !exists {
			m[nums[i]] = struct{}{}
		}
	}

	for {
		if _, exists := m[original]; exists {
            original *= 2
            continue
		}
        break
	}

	return original
}

// array, hash map
// time: O(n + log(M))
// space: O(n)