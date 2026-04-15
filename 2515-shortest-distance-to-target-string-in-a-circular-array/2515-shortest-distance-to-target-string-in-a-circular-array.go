func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	found := false
    rd, ld := 0, 0

	// check next
	i := startIndex
	for {
		if words[i] == target {
			found = true
			break
		}

		i = (i + 1) % n
        rd++
		if i == startIndex {
			break
		}
	}

	if !found {
		return -1
	}

	// check prev
	i = startIndex
	for {
		if words[i] == target {
			break
		}

		i = (i - 1 + n) % n
        ld++
		if i == startIndex {
			break
		}
	}

	return min(rd, ld)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// array, string
// time: O(n)
// space: O(1)