func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var sum int64 = 0
	var minExtra int64 = math.MaxInt64
	count := 0

	for _, val := range nums {
		if (val ^ k) > val {
			sum += int64(val ^ k)
			minExtra = min(minExtra, int64((val^k)-val))
			count++
		} else {
			sum += int64(val)
			minExtra = min(minExtra, int64(val-(val^k)))
		}
	}

	if count%2 == 0 {
		return sum
	}

	return sum - minExtra
}

// NOTICE ME SENPAI!!!