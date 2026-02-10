func longestBalanced(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		evenSet, oddSet := map[int]bool{}, map[int]bool{}
		for j := i; j < len(nums); j++ {
			if nums[j]%2 == 0 { // even
				evenSet[nums[j]] = true
			} else { // odd
				oddSet[nums[j]] = false
			}

			if len(evenSet) == len(oddSet) {
				ans = max(ans, j-i+1)
			}
		}
	}

	return ans
}

// array, hash map
// time: O(n^2)
// space: O(n)