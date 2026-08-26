func shortestBeautifulSubstring(s string, k int) string {
	ss := ""
	count := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			count++
		}

		if count == k {
			for s[left] == '0' {
				left++
			}

			curr := s[left : right+1]
			if ss == "" || len(curr) < len(ss) || (len(curr) == len(ss)) && (curr < ss) {
				ss = curr
			}

            // discard the first '1' so that the next window looks for the next k '1's
            left++
            count--
		}
	}

	return ss
}

// sliding window, string
// time: O(n)
// space: O(1)
