func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	rec := nums[0]
	for _, num := range nums {
		if num-rec > k {
			ans++
			rec = num
		}
	}
	return ans
}

// NOTICE ME SENPAI!!!