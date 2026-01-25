func minimumPairRemoval(nums []int) int {
	countDecrease := 0
	ops := 0
	h := &MinHeap{}
	merged := make([]bool, len(nums))
	head := &Node{val: int64(nums[0]), left: 0}
	curr := head
    // translate the nums into doubly-linked list and push pairs into heap
	for i := 1; i < len(nums); i++ {
		newNode := &Node{val: int64(nums[i]), left: i}
		curr.next = newNode
		newNode.prev = curr

		heap.Push(h, &Pair{
			first:  curr,
			second: newNode,
			cost:   curr.val + newNode.val,
		})

        // count decreasing adjacent number
		if nums[i-1] > nums[i] {
			countDecrease++
		}

		curr = newNode
	}

	for countDecrease > 0 {
		pair := heap.Pop(h).(*Pair)
		first := pair.first
		second := pair.second
		cost := pair.cost

		if merged[first.left] || merged[second.left] || first.val+second.val != cost {
			continue
		}

		ops++ // do the operation
		if first.val > second.val {
			countDecrease--
		}

		prevNode := first.prev
		nextNode := second.next
		first.next = nextNode // remove second
		if nextNode != nil {
			nextNode.prev = first
		}

		// prevNode and first (first)
		if prevNode != nil {
			// before and after condition
			if prevNode.val > first.val && prevNode.val <= cost {
				countDecrease--
			} else if prevNode.val <= first.val && prevNode.val > cost {
				countDecrease++
			}
			heap.Push(h, &Pair{
				first:  prevNode,
				second: first,
				cost:   prevNode.val + cost,
			})
		}

		// first (second) abd nextNode
		if nextNode != nil {
			// before and after condiition
			if second.val > nextNode.val && cost <= nextNode.val {
				countDecrease--
			} else if second.val <= nextNode.val && cost > nextNode.val {
				countDecrease++
			}
			heap.Push(h, &Pair{
				first:  first,
				second: nextNode,
				cost:   cost + nextNode.val,
			})
		}

        first.val = cost
        merged[second.left] = true
	}

    return ops
}

// array, doubly-linked list, heap, linked list, simulation
// time: O(nlog(n))
// space: O(n)

type Node struct {
	val  int64
	left int
	prev *Node
	next *Node
}

type Pair struct {
	first  *Node
	second *Node
	cost   int64
}

type MinHeap []*Pair

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].first.left < h[j].first.left
	}
	return h[i].cost < h[j].cost
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(*Pair)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	pair := old[n-1]
	*h = old[:n-1]
	return pair
}