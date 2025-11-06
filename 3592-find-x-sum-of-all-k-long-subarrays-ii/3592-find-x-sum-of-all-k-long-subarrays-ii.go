func findXSum(nums []int, k int, x int) []int64 {
	// smallHeap optimization: if x == k -> sum of window
	if x == k {
		ans := []int64{}
		windowSum := 0
		for i := 0; i < len(nums); i++ {
			windowSum += nums[i]
			if i >= k {
				windowSum -= nums[i-k]
			}
			if i >= k-1 {
				ans = append(ans, int64(windowSum))
			}
		}
		return ans
	}

	helper := NewHelper(x)
	ans := []int64{}
	for i := 0; i < len(nums); i++ {
		helper.Insert(nums[i])
		if i >= k {
			helper.Remove(nums[i-k])
		}
		if i >= k-1 {
			ans = append(ans, helper.Get())
		}
	}
	return ans
}

// array, hash map, heap, sliding window
// time: O(n log(n))
// space: O(k)

// =================
// UTILS
// =================
func compare(a, b pair) int {
	if a.freq == b.freq {
		return a.num - b.num
	}
	return a.freq - b.freq
}

// =================
// DATA STRUCTURES
// =================
type pair struct {
	num  int
	freq int
}

// MinHeap (num, freq)
type MinHeap []pair

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Less(i, j int) bool { return compare((*h)[i], (*h)[j]) < 0 }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MaxHeap (num, freq)
type MaxHeap []pair

func (h *MaxHeap) Len() int { return len(*h) }

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Less(i, j int) bool { return compare((*h)[i], (*h)[j]) > 0 }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Helper struct with 2 heaps + maps
type Helper struct {
	x          int
	result     int64        // sum of num * freq for elements currently in largeHeap
	occ        map[int]int  // occur; map (num -> freq) of current numbers in window
	inLarge    map[int]bool // set (map) of current numbers in largeHeap
	largeHeap  *MinHeap     // MinHeap; top is the superior pair in largeHeap
	smallHeap  *MaxHeap     // MaxHeap; top is the superior pair in smallHeap
	largeCount int          // number of distinct elements in largeHeap
}

func NewHelper(x int) *Helper {
	return &Helper{
		x:         x,
		result:    0,
		occ:       make(map[int]int),
		inLarge:   make(map[int]bool),
		largeHeap: &MinHeap{},
		smallHeap: &MaxHeap{},
	}
}

func (h *Helper) cleanLargeTop() (pair, bool) {
	for h.largeHeap.Len() > 0 {
		top := (*h.largeHeap)[0]
		currFreq, ok := h.occ[top.num]
		if !ok || currFreq != top.freq || !h.inLarge[top.num] {
			heap.Pop(h.largeHeap)
			continue
		}
		return top, true
	}
	return pair{}, false
}

func (h *Helper) cleanSmallTop() (pair, bool) {
	for h.smallHeap.Len() > 0 {
		top := (*h.smallHeap)[0]
		currFreq, ok := h.occ[top.num]
		if !ok || currFreq != top.freq || h.inLarge[top.num] {
			heap.Pop(h.smallHeap)
			continue
		}
		return top, true
	}
	return pair{}, false
}

func (h *Helper) promoteFromSmall() bool {
	p, ok := h.cleanSmallTop()
	if !ok {
		return false
	}

	heap.Pop(h.smallHeap)

	h.inLarge[p.num] = true
	h.largeCount++
	h.result += int64(p.num) * int64(p.freq)
	heap.Push(h.largeHeap, p)
	return true
}

func (h *Helper) demoteFromLarge() bool {
	p, ok := h.cleanLargeTop()
	if !ok {
		return false
	}

	heap.Pop(h.largeHeap)

	h.inLarge[p.num] = false
	h.largeCount--
	h.result -= int64(p.num) * int64(p.freq)
	heap.Push(h.smallHeap, p)
	return true
}

// ensure invariants:
// - largeCount <= x
// - if largeCount < x, promote from smallHeap until largeCount == min(x, largeCount)
// - if there is a pair in smallHeap "more superior" than minLarge, swap
func (h *Helper) rebalance() {
	// promote until largeCount == min(x, largeCount)
	for h.largeCount < h.x {
		if !h.promoteFromSmall() {
			break
		}
	}

	// if maxSmall is "more superior" than minLarge, swap (until achieving equilibrium)
	for {
		minLarge, okL := h.cleanLargeTop()
		maxSmall, okS := h.cleanSmallTop()
		if !okL || !okS {
			break
		}
		if compare(maxSmall, minLarge) > 0 {
			// demote minLarge
			heap.Pop(h.largeHeap)
			h.inLarge[minLarge.num] = false
			h.largeCount--
			h.result -= int64(minLarge.num) * int64(minLarge.freq)
			heap.Push(h.smallHeap, minLarge)

			// promote maxSmall
			heap.Pop(h.smallHeap)
			h.inLarge[maxSmall.num] = true
			h.largeCount++
			h.result += int64(maxSmall.num) * int64(maxSmall.freq)
			heap.Push(h.largeHeap, maxSmall)
		} else {
			break
		}
	}
}

// insert a number (increment its freq)
func (h *Helper) Insert(num int) {
	oldFreq := h.occ[num]
	// if num existed and was in largeHeap, deduct its old contribution; re-insert new freq
	if oldFreq > 0 && h.inLarge[num] {
		h.result -= int64(num) * int64(oldFreq)
		h.inLarge[num] = false
		h.largeCount--
		// NOTE: stale entry remains in largeHeap, but will be lazily deleted later
	}
	h.occ[num] = oldFreq + 1
	newPair := pair{num: num, freq: oldFreq + 1}

	if h.largeCount < h.x {
		h.inLarge[num] = true
		h.largeCount++
		h.result += int64(newPair.num) * int64(newPair.freq)
		heap.Push(h.largeHeap, newPair)
	} else {
		// compare with minLarge
		minLarge, ok := h.cleanLargeTop()
		if !ok {
			// no valid minLarge (largeHeap was empty) -> put to largeHeap
			h.inLarge[num] = true
			h.largeCount++
			h.result += int64(newPair.num) * int64(newPair.freq)
			heap.Push(h.largeHeap, newPair)
		} else {
			// if newPair is "more superior" than minLarge -> demote minLarge, put newPair into largeHeap
			if compare(newPair, minLarge) > 0 {
				// demote minLarge
				heap.Pop(h.largeHeap)
				h.inLarge[minLarge.num] = false
				h.largeCount--
				h.result -= int64(minLarge.num) * int64(minLarge.freq)
				heap.Push(h.smallHeap, minLarge)

				// insert newPair into largeHeap
				h.inLarge[newPair.num] = true
				h.largeCount++
				h.result += int64(newPair.num) * int64(newPair.freq)
				heap.Push(h.largeHeap, newPair)
			} else {
				// put newPair into smallHeap
				h.inLarge[newPair.num] = false
				heap.Push(h.smallHeap, newPair)
			}
		}
	}
	h.rebalance()
}

// remove a number (decrement its freq)
func (h *Helper) Remove(num int) {
	oldFreq := h.occ[num]
	if oldFreq == 0 {
		return // shouldn't happen
	}
	// if currently in largeHeap, deduct its old contribution immediately
	if h.inLarge[num] {
		h.result -= int64(num) * int64(oldFreq)
		h.inLarge[num] = false
		h.largeCount--
		// NOTE: stale entry remains in largeHeap, but will be lazily deleted later
	}
	h.occ[num] = oldFreq - 1
	if h.occ[num] == 0 {
		delete(h.occ, num)
	} else {
        // insert decreased freq version
		newPair := pair{num: num, freq: oldFreq - 1}
		// if largeHeap has space, put into largeHeap; else, compare with minLarge and place accordingly
		if h.largeCount < h.x {
			h.inLarge[num] = true
			h.largeCount++
			h.result += int64(newPair.num) * int64(newPair.freq)
			heap.Push(h.largeHeap, newPair)
		} else {
			minLarge, ok := h.cleanLargeTop()
			if !ok {
				h.inLarge[num] = true
				h.largeCount++
				h.result += int64(newPair.num) * int64(newPair.freq)
				heap.Push(h.largeHeap, newPair)
			} else {
				if compare(newPair, minLarge) > 0 {
					// demote minLarge
					heap.Pop(h.largeHeap)
					h.inLarge[minLarge.num] = false
					h.largeCount--
					h.result -= int64(minLarge.num) * int64(minLarge.freq)
					heap.Push(h.smallHeap, minLarge)

					// insert newPair into largeHeap
					h.inLarge[newPair.num] = true
					h.largeCount++
					h.result += int64(newPair.num) * int64(newPair.freq)
					heap.Push(h.largeHeap, newPair)
				} else {
					// put newPair into smallHeap
					h.inLarge[newPair.num] = false
					heap.Push(h.smallHeap, newPair)
				}
			}
		}
	}
	h.rebalance()
}

func (h *Helper) Get() int64 {
	return h.result
}
