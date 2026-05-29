func numberOfSpecialChars(word string) int {
	ans := 0
	m := map[rune]bool{}
	for _, c := range word {
		if 'a' <= c && c <= 'z' {
			m[c-32] = true
		}
	}

	for _, c := range word {
		if m[c] {
			ans++
			delete(m, c)
		}
	}

	return ans
}

// hash map
// time: O(n)
// space: O(k)