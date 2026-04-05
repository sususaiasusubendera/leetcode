func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		// vertical
		case 'U':
			y++
		case 'D':
			y--
		// horizontal
		case 'R':
			x++
		case 'L':
            x--
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return false
}

// string
// time: O(n)
// space: O(1)