func leftRightDifference(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	leftSum, rightSum := 0, 0
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
	}

	for i := 0; i < n; i++ {
		ans[i] = abs(leftSum - rightSum)
		leftSum += nums[i]
		if i < n-1 {
			rightSum -= nums[i+1]
		}
	}

	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// array, prefix sum
// time: O(n)
// space: O(1)