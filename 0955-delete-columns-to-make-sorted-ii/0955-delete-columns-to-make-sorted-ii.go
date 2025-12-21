func minDeletionSize(strs []string) int {
	deletions := 0
	skipWord := make(map[int]bool) // when we make sure a col is good, why would we even check that word again?!?
outerloop:
	for colIdx := 0; colIdx < len(strs[0]); colIdx++ {
		// check that the current column does not break the order for any word
		toSkipIfWeKeepThisColumn := map[int]bool{}
		for wordIdx := 1; wordIdx < len(strs); wordIdx++ {
			if !skipWord[wordIdx] && strs[wordIdx][colIdx] < strs[wordIdx-1][colIdx] {
				deletions++
				continue outerloop
			} else if strs[wordIdx][colIdx] > strs[wordIdx-1][colIdx] {
				toSkipIfWeKeepThisColumn[wordIdx] = true
			}
		}

		// now we are sure current col is good, so for every str where str[col] > prevStr[col] order is LOCKED and we can skip it
		for col := range toSkipIfWeKeepThisColumn {
			skipWord[col] = true
		}
	}
	return deletions
}

// NOTICE ME SENPAI
// solution from solutions @firas5445
// array, string, greedy
