func minimumDistance(nums []int) int {
    n := len(nums)
    if n < 3 {
        return -1
    }

	x := int(10e10)
    exists := false
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] == nums[j] && nums[j] == nums[k] && nums[i] == nums[k] {
					x = min(x, abs(i-j)+abs(j-k)+abs(k-i))
                    exists = true
				}
			}
		}
	}

	if exists {
        return x
    }
    return -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// array
// time: O(n^3)
// space: O(1)