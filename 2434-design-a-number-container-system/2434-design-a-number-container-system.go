// MAIN
type NumberContainers struct {
    mapIdxNum map[int]int // index -> number
    mapNumIdxs map[int]*MinHeap // number -> indexes (priority queue)
    validIdx map[int]map[int]bool // number -> (indexes -> bool (set of indexes))
}


func Constructor() NumberContainers {
    return NumberContainers {
        mapIdxNum: map[int]int{},
        mapNumIdxs: map[int]*MinHeap{},
        validIdx: map[int]map[int]bool{},
    }
}


func (this *NumberContainers) Change(index int, number int) {
    // delete the index from validIdx[oldNum] (the index becomes invalid) if it already exists
    if oldNum, exists := this.mapIdxNum[index]; exists {
        delete(this.validIdx[oldNum], index)
    }

    // update index with the latest number
    this.mapIdxNum[index] = number

    // initialize a heap (priority queue) if it's a new number
    if _, exists := this.mapNumIdxs[number]; !exists {
        this.mapNumIdxs[number] = &MinHeap{}
        heap.Init(this.mapNumIdxs[number])
        this.validIdx[number] = map[int]bool{}
    }

    // push the index into the priority queue and mark it valid (in validIdx[number])
    heap.Push(this.mapNumIdxs[number], index)
    this.validIdx[number][index] = true
}


func (this *NumberContainers) Find(number int) int {
    // return -1 if the number doesn't exist
    if _, exists := this.mapNumIdxs[number]; !exists {
        return -1
    }

    // get the smallest index that is still valid 
    for this.mapNumIdxs[number].Len() > 0 {
        top := (*this.mapNumIdxs[number])[0]
        if this.validIdx[number][top] {
            return top
        }
        heap.Pop(this.mapNumIdxs[number])
    }

    return -1
}

// map and priority queue approach
// time: O(log k + m log k)
// space: O(n + mk)

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
**/

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

// -------

// >>> authentic TLE map solution <<<

// type NumberContainers struct {
//     mapIdxNum map[int]int
//     mapNumIdxs map[int][]int
// }


// func Constructor() NumberContainers {
//     return NumberContainers{
//         map[int]int{},
//         map[int][]int{},
//     }
// }


// func (this *NumberContainers) Change(index int, number int)  {
//     if oldNum, exists := this.mapIdxNum[index]; exists {
//         idxs := this.mapNumIdxs[oldNum]
//         for i, idx := range idxs {
//             if idx == index {
//                 this.mapNumIdxs[oldNum] = append(idxs[:i], idxs[i+1:]...)
//             }
//             break
//         }

//         if len(this.mapNumIdxs[oldNum]) == 0 {
//             delete(this.mapNumIdxs, oldNum)
//         }
//     }

//     this.mapIdxNum[index] = number
//     this.mapNumIdxs[number] = append(this.mapNumIdxs[number], index)
//     sort.Ints(this.mapNumIdxs[number])
// }


// func (this *NumberContainers) Find(number int) int {
//     if len(this.mapNumIdxs[number]) == 0 {
//         return -1
//     }
//     return this.mapNumIdxs[number][0]
// }