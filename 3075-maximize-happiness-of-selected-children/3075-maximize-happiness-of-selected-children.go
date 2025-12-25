func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool { return happiness[i] > happiness[j] })

	sub := 0
	idx := 0
	ans := 0
	for k > 0 {
		if happiness[idx]-sub <= 0 {
			break
		}

		ans += happiness[idx] - sub

		k--
		sub++
		idx++
	}

	return int64(ans)
}

// array, greedy, sorting
// time: O(nlog(n))
// space: O(1)