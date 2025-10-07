func avoidFlood(rains []int) []int {
	nextRain := map[int][]int{}
	for i, lake := range rains {
		if lake > 0 {
			nextRain[lake] = append(nextRain[lake], i)
		}
	}

	full := map[int]struct{}{}
	h := &MinHeap{}
	ans := []int{}
	for i := 0; i < len(rains); i++ {
		// rain at lake rains[i] today
		if rains[i] > 0 {
			lake := rains[i]

			if _, ok := full[lake]; ok { return []int{} } // flood occurs

			full[lake] = struct{}{}
			ans = append(ans, -1)
			if len(nextRain[lake]) > 0 {
				nextRain[lake] = nextRain[lake][1:]
				if len(nextRain[lake]) > 0 {
					heap.Push(h, Rain{lake, nextRain[lake][0]})
				}
			}

			continue
		}

		// no rain today
		if (*h).Len() > 0 {
			nr := heap.Pop(h).(Rain) // nr: next rain
			if _, ok := full[nr.lake]; ok {
				delete(full, nr.lake)
				ans = append(ans, nr.lake)
			} else {
				ans = append(ans, 1)
			}
		} else {
			ans = append(ans, 1)
		}
	}

	return ans
}

// array, hash map, heap
// time: O(n log(n))
// space: O(n)

type Rain struct {
	lake int
	day  int
}

// min heap implementation
type MinHeap []Rain

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Less(i, j int) bool { return (*h)[i].day < (*h)[j].day }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(Rain)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}