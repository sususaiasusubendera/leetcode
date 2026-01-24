func minPairSum(nums []int) int {
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ans := 0
	left, right := 0, len(nums)-1
	for left < right {
		ans = max(ans, nums[left] + nums[right])
		left++
		right--
	}

	return ans
}

// array, greedy, sorting, two pointers
// time: O(nlog(n))
// space: O(1)