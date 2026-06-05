func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool { return cost[i] > cost[j] })

	idx := 0
	total := 0
	for idx <= len(cost)-1 {
		if idx <= len(cost)-3 {
			total += cost[idx] + cost[idx+1]
			idx += 3
		} else if idx == len(cost)-2 {
			total += cost[idx] + cost[idx+1]
            idx += 2
		} else if idx == len(cost)-1 {
			total += cost[idx]
            idx++
		}
	}

	return total
}

// array, greedy, sorting
// time: O(nlog(n))
// space: O(1)