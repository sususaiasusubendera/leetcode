func minimumIndex(nums []int) int {
	freq := make(map[int]int)
	dom := -1
	for _, num := range nums {
		freq[num]++
		if freq[num] > freq[dom] {
			dom = num
		}
	}

	b := 0
	for i := range nums {
		if nums[i] == dom {
			b++
		}
		before := i+1 < b*2
		after := (len(nums) - i - 1) < (freq[dom]-b)*2
		if before && after {
			return i
		}
	}

	return -1
}

// notice me senpai!!!