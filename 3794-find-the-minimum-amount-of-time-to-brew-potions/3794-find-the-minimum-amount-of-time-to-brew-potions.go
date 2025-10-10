func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	wizardTime := make([]int, len(skill)) // store the finish time of wizard[i], used as the next potion's start time
	for j := 0; j < m; j++ {
		currTime := 0
		for i := 0; i < n; i++ {
			if currTime < wizardTime[i] {
				currTime = wizardTime[i]
			} 

			currTime += skill[i] * mana[j]
		}

		wizardTime[n-1] = currTime
		for i := n - 2; i >= 0; i-- {
			wizardTime[i] = wizardTime[i+1] - (skill[i+1] * mana[j])
		}
	}

	return int64(wizardTime[n-1])
}

// notice me senpai