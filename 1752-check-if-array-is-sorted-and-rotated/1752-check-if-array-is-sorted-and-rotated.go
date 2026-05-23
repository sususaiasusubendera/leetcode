func check(nums []int) bool {
	sorted := true
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			count++
		}
	}

	if sorted {
		return true
	} else {
		if count > 1 {
			return false
		} else {
			if nums[0] >= nums[len(nums)-1] {
				return true
			} else {
				return false
			}
		}
	}
}

// array
// time: O(n)
// space: O(1)

// #1 2025/02/02
// #2 2026/05/23