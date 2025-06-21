func minimumDeletions(word string, k int) int {
	cnt := make(map[rune]int)
	for _, ch := range word {
		cnt[ch]++
	}
	res := len(word)
	for _, a := range cnt {
		deleted := 0
		for _, b := range cnt {
			if a > b {
				deleted += b
			} else if b > a+k {
				deleted += b - (a + k)
			}
		}
		if deleted < res {
			res = deleted
		}
	}
	return res
}

// editorial
// NOTICE ME SENPAI!!!