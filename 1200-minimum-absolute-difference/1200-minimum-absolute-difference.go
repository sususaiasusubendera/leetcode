func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	minDiff := (1 << 31) - 1
	for i := 1; i < len(arr); i++ {
		minDiff = min(minDiff, arr[i]-arr[i-1])
	}

	ans := [][]int{}
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}

// array, sorting
// time: O(nlog(n))
// space: O(n)