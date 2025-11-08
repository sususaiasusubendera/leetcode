func maxPower(stations []int, r int, k int) int64 {
	// compute the base power for each city (total power currently received)
	base := make([]int, len(stations))
	prefix := make([]int, len(stations)+1)
	for i := 0; i < len(stations); i++ {
		prefix[i+1] = prefix[i] + stations[i]
	}
	for i := 0; i < len(stations); i++ {
		left := max(0, i-r)
		right := min(len(stations)-1, i+r)
		base[i] = prefix[right+1] - prefix[left]
	}

	// binary search
	low := 0
	total := 0
	for _, station := range stations {
		total += station
	}
	high := total + k
	res := 0
	for low <= high {
		mid := (low + high) / 2
		if canAchieve(base, r, k, mid) {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return int64(res)
}

// array, binary search, diff arr, greedy, prefix sum
// time: O(n.log(sum(stations)+k))
// space: O(n)

func canAchieve(base []int, r int, k int, target int) bool {
	diff := make([]int, len(base)+1)
	currAdd := 0
	used := 0
	for i := 0; i < len(base); i++ {
		currAdd += diff[i]        // current active diffArr effect
		curr := base[i] + currAdd // current actual power of city i
		if curr < target {
			need := target - curr
			used += need
			if used > k {
				return false
			}

			pos := min(i+r, len(base)-1)
			left := max(0, pos-r)
			right := min(len(base)-1, pos+r)

			diff[left] += need
			diff[right+1] -= need
			currAdd += need
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

