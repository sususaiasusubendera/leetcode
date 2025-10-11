func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	wizardTime := make([]int, len(skill)) // stores the time when each wizard can start brewing the next potion
	for j := 0; j < m; j++ {
		currTime := 0
		for i := 0; i < n; i++ {
            // wait until this wizard is available
			if currTime < wizardTime[i] {
				currTime = wizardTime[i]
			} 

			currTime += skill[i] * mana[j]
		}

        // backtrack to update when each wizard can start the next potion
		wizardTime[n-1] = currTime
		for i := n - 2; i >= 0; i-- {
			wizardTime[i] = wizardTime[i+1] - (skill[i+1] * mana[j])
		}
	}

	return int64(wizardTime[n-1])
}

// array, dp(?), simulation
// time: O(mn)
// space: O(n)