func maximumInvitations(favorite []int) int {
	n := len(favorite)
	inDegree := make([]int, n)
	for _, f := range favorite {
		inDegree[f]++
	}

	queue := []int{}
	for i := range n {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	longestChain := make([]int, n)
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		fav := favorite[i]
		longestChain[fav] = max(longestChain[fav], longestChain[i]+1)
		inDegree[fav]--
		if inDegree[fav] == 0 {
			queue = append(queue, fav)
		}
	}

	maxCycleSize := 0
	mutualPairs := 0
	for i := range n{
		if inDegree[i] == 0 {
			continue
		}
		cycleSize := 0
		cur := i
		for inDegree[cur] > 0 {
			inDegree[cur] = 0
			cur = favorite[cur]
			cycleSize++
		}
		if cycleSize == 2 {
			mutualPairs += 2 + longestChain[i] + longestChain[favorite[i]]
		} else {
			maxCycleSize = max(maxCycleSize, cycleSize)
        }
	}

	return max(maxCycleSize, mutualPairs)
}

// notice me senpai