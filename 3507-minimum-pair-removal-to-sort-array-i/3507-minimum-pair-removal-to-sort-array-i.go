func minimumPairRemoval(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	ops := 0
	for len(nums) > 1 {
		sorted := true
		minSum := 1_000_000_007
		targetIdx := -1 // idx delete

		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				sorted = false
			}

			sum := nums[i] + nums[i-1]
			if sum < minSum {
				minSum = sum
				targetIdx = i - 1
			}
		}

		if sorted {
			break
		}

		if targetIdx > -1 {
			ops++
			nums[targetIdx] = minSum
			nums = append(nums[:targetIdx+1], nums[targetIdx+2:]...)
		}
	}

	return ops
}

// array, simulation
// time: O(n^2)
// space: O(n)