func checkCuts(rectangles [][]int, dim int) bool {
	gapCount := 0

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][dim] < rectangles[j][dim]
	})

	furthestEnd := rectangles[0][dim+2]

	for i := 1; i < len(rectangles); i++ {
		rect := rectangles[i]

		if furthestEnd <= rect[dim] {
			gapCount++
		}

		if rect[dim+2] > furthestEnd {
			furthestEnd = rect[dim+2]
		}
	}

	return gapCount >= 2
}

// NOTICE ME SENPAI!!!

func checkValidCuts(n int, rectangles [][]int) bool {
	return checkCuts(rectangles, 0) || checkCuts(rectangles, 1)
}
