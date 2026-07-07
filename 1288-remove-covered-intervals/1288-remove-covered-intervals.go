func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	idx := 1
    for idx < len(intervals) {
        if intervals[idx-1][1] >= intervals[idx][1] {
            intervals = append(intervals[:idx], intervals[idx+1:]...)
        } else {
            idx++
        }
    }

    return len(intervals)
}

// array, sorting
// time: O(n^2)
// space: O(1)