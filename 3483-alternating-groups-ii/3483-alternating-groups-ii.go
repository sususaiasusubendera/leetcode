func numberOfAlternatingGroups(colors []int, k int) int {
	extendedColors := append(colors, colors[:k-1]...)

	n := len(extendedColors)
	countAltGroups := 0
	left, right := 0, 1

	for right < n {
		if extendedColors[right] == extendedColors[right-1] {
			left = right
			right++
			continue
		}

		right++

		if right-left < k {
			continue
		}

		countAltGroups++
		left++
	}

    return countAltGroups
}

// O(n)
// O(n)