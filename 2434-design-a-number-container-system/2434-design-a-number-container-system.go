type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Ascending order
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Main structure

type NumberContainers struct {
    mapIdxNum  map[int]int   // {index: number}
    mapNumIdxs map[int]*MinHeap // {number: heap of indices}
}

func Constructor() NumberContainers {
    return NumberContainers{
        mapIdxNum:  make(map[int]int),
        mapNumIdxs: make(map[int]*MinHeap),
    }
}

func (this *NumberContainers) Change(index int, number int) {
    // If the index already has an existing number, remove it from the previous heap
    if oldNum, exists := this.mapIdxNum[index]; exists && oldNum != number {
        // No need to remove it directly from the heap, just mark that this index no longer belongs to `oldNum`
        heap.Init(this.mapNumIdxs[oldNum]) // Rebuild heap to clean up invalid indices
    }

    // Update the number at the index
    this.mapIdxNum[index] = number

    // Insert index into the heap for the new number
    if this.mapNumIdxs[number] == nil {
        this.mapNumIdxs[number] = &MinHeap{}
    }
    heap.Push(this.mapNumIdxs[number], index)
}

func (this *NumberContainers) Find(number int) int {
    if h, exists := this.mapNumIdxs[number]; exists {
        // Ensure the heap is still valid with the correct indices
        for h.Len() > 0 {
            smallestIndex := (*h)[0] // Check the smallest index
            if this.mapIdxNum[smallestIndex] == number {
                return smallestIndex
            }
            heap.Pop(h) // Remove outdated index
        }
    }
    return -1
}