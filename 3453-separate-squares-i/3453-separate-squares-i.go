func separateSquares(squares [][]int) float64 {
	// find max y and total area
	maxY, totalArea := 0.0, 0.0
	for _, s := range squares {
		y, l := s[1], s[2]
		totalArea += float64(l * l)
		if float64(y+l) > maxY {
			maxY = float64(y + l)
		}
	}

	var check func(limitY float64) bool
	check = func(limitY float64) bool {
		area := 0.0
		for _, s := range squares {
			y, l := s[1], s[2]
			if float64(y) < limitY {
                // cut off by limit y or not (l = y)
				overlap := math.Min(limitY-float64(y), float64(l))
				area += float64(l) * overlap
			}
		}

		return area >= totalArea/2.0
	}

    // binary search
    lo, hi := 0.0, maxY
    eps := 1e-5
    for math.Abs(hi - lo) > eps {
        mid := (hi + lo) / 2.0
        if check(mid) {
            hi = mid
        } else {
            lo = mid
        }
    }

    return hi
}

// editorial
// array, binary search
// time: O(nlog(maxY/eps))
// space: O(1)