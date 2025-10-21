func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	ans := 0
	numCount := make(map[int]int)

	lastNumIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[lastNumIndex] {
			numCount[nums[lastNumIndex]] = i - lastNumIndex
			ans = max(ans, i-lastNumIndex)
			lastNumIndex = i
		}
	}

	numCount[nums[lastNumIndex]] = len(nums) - lastNumIndex
	ans = max(ans, len(nums)-lastNumIndex)

	leftBound := func(value int) int {
		left, right := 0, len(nums)-1
		for left < right {
			mid := (left + right) / 2
			if nums[mid] < value {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	rightBound := func(value int) int {
		left, right := 0, len(nums)-1
		for left < right {
			mid := (left + right + 1) / 2
			if nums[mid] > value {
				right = mid - 1
			} else {
				left = mid
			}
		}
		return left
	}

	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		l := leftBound(i - k)
		r := rightBound(i + k)

		tempAns := 0
		if count, exists := numCount[i]; exists {
			tempAns = min(r-l+1, count+numOperations)
		} else {
			tempAns = min(r-l+1, numOperations)
		}
		ans = max(ans, tempAns)
	}

	return ans
}

// EDITORIAL
// NOTICE ME SENPAI
// sort + enumerate + bin search