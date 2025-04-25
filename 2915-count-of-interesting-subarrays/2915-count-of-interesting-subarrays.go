func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	var res int64 = 0
	var prefix int = 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefix++
		}
		res += int64(cnt[(prefix-k+modulo)%modulo])
		cnt[prefix%modulo]++
	}
	return res
}

// notice me senpai