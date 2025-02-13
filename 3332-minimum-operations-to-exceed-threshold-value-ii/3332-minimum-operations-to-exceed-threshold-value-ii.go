import "container/heap"

func minOperations(nums []int, k int) int {
    if len(nums) < 2 {
        return 0
    }

    // initialize a priority queue (pq)
    pq := &MinHeap{}
    heap.Init(pq)
    for i := 0; i < len(nums); i++ { // add all elements in nums to pq
        heap.Push(pq, nums[i])
    }

    // do the operation until all elements in pq are greater than or equal to k
    numOps := 0
    for pq.Len() >= 2 && ((*pq)[0] < k || (*pq)[1] < k) {
        num1 := heap.Pop(pq).(int)
        num2 := heap.Pop(pq).(int)
        newNum := (2 * min(num1, num2)) + max(num1, num2)
        numOps++
        heap.Push(pq, newNum)
    }

    return numOps
}

// priority queue approach
// time: O(n.log(n))
// space: O(n)

// MINHEAP
type MinHeap []int

func (h *MinHeap) Len() int {
    return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
    return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
    (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// UTILS
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}