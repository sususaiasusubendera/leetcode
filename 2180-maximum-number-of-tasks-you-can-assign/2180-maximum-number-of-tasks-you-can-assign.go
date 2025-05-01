func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	check := func(mid int) bool {
		p := pills
		ws := list.New() // Use deque
		ptr := m - 1
		// Enumerate each task from largest to smallest
		for i := mid - 1; i >= 0; i-- {
			for ptr >= m-mid && workers[ptr]+strength >= tasks[i] {
				ws.PushFront(workers[ptr]) // Add to the front of the queue
				ptr--
			}
			if ws.Len() == 0 {
				return false
			}
			// If the largest element in the deque is greater than or equal to
			// tasks[i]
			if ws.Back().Value.(int) >= tasks[i] {
				ws.Remove(ws.Back()) // Remove from the end of the queue
			} else {
				if p == 0 {
					return false
				}
				p--
				ws.Remove(ws.Front()) // Remove the front of the queue
			}
		}
		return true
	}

	left, right, ans := 1, min(m, n), 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// (editorial) notice me senpai