func decodeCiphertext(encodedText string, rows int) string {
	m, n := rows, len(encodedText)/rows

	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	r, c := 0, 0
	for i := 0; i < len(encodedText); i++ {
		grid[r][c] = encodedText[i]
		c++
		if c == n {
			c = 0
			r++
		}
	}

	ans := []byte{}
	x := 0
	r, c = 0, 0
	for x < n {
		ans = append(ans, grid[r][c])
		r++
		c++
		if r == m || c == n {
			r = 0
			c = x + 1
			x++
		}
	}

	idx := len(ans) - 1
	for idx >= 0 && ans[idx] == ' ' {
		idx--
	}

	return string(ans[:idx+1])
}

// string
// time: O(n)
// space: O(n)