func minSum(nums1 []int, nums2 []int) int64 {
	var sum1, sum2 int64
	var zero1, zero2 int

	for _, num := range nums1 {
		sum1 += int64(num)
		if num == 0 {
			sum1++
			zero1++
		}
	}

	for _, num := range nums2 {
		sum2 += int64(num)
		if num == 0 {
			sum2++
			zero2++
		}
	}

	if (zero1 == 0 && sum2 > sum1) || (zero2 == 0 && sum1 > sum2) {
		return -1
	}

	if sum1 > sum2 {
		return sum1
	}
	return sum2
}

// editorial
// notice me senpai