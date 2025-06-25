func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n1 := len(nums1)
	left, right := int64(-1e10), int64(1e10)
	for left <= right {
		mid := (left + right) / 2
		count := int64(0)
		for i := 0; i < n1; i++ {
			count += int64(f(nums2, int64(nums1[i]), mid))
		}
		if count < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// binary search
// editorial
// hard, skipped
// NOTICE ME SENPAI!

func f(nums2 []int, x1 int64, v int64) int {
	n2 := len(nums2)
	left, right := 0, n2-1
	for left <= right {
		mid := (left + right) / 2
		prod := int64(nums2[mid]) * x1
		if (x1 >= 0 && prod <= v) || (x1 < 0 && prod > v) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if x1 >= 0 {
		return left
	} else {
		return n2 - left
	}
}