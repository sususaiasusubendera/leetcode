func maxActiveSectionsAfterTrade(s string) int {
	count1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count1++
		}
	}

	zeroBlocks := []int{}
	i := 0
	for i < len(s) {
		start := i
		for i < len(s) && s[i] == s[start] {
			i++
		}
		if s[start] == '0' {
			zeroBlocks = append(zeroBlocks, i-start)
		}
	}

	if len(zeroBlocks) < 2 {
		return count1
	}

	best := 0
	for j := 1; j < len(zeroBlocks); j++ {
		best = max(best, zeroBlocks[j]+zeroBlocks[j-1])
	}

	return count1 + best
}

// greedy
// time: O(n)
// space: O(n)