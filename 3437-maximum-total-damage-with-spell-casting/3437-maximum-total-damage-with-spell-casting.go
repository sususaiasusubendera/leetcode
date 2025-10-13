func maximumTotalDamage(power []int) int64 {
	freq := map[int]int{}
	for i := 0; i < len(power); i++ {
		freq[power[i]] += power[i]
	}

	keys := []int{}
	for key := range freq {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	dp := make([]int, len(keys))
	dp[0] = freq[keys[0]]
	for i := 1; i < len(keys); i++ {
		dp[i] = dp[i-1]

		take := freq[keys[i]]

		j := i - 1
		for j >= 0 && keys[i]-keys[j] <= 2 {
			j--
		}

		if j >= 0 {
			take += dp[j]
		}

		if take > dp[i] {
			dp[i] = take
		}
	}

	return int64(dp[len(dp)-1])
}

// array, dp (bot-up), hash map, sorting
// time: O(n log(n))
// space: O(n)