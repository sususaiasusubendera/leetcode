func numberOfPairs(points [][]int) int {
	count := 0
	slices.SortFunc(points, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return b[1] - a[1]
	})
	for i := 0; i < len(points); i++ {
		a := points[i]
		xMin, xMax := a[0] - 1, int(1e10)
		yMin, yMax := int(-1e10), a[1] + 1
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			if b[0] > xMin && b[0] < xMax && b[1] > yMin && b[1] < yMax {
				count++
				xMin = b[0]
                yMin = b[1]
			}
        }
	}

	return count
}

// time: O(n^2)
// space: O(n)