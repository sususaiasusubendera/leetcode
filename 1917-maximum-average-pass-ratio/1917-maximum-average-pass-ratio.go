func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    h := &maxHeap{}
    heap.Init(h)
    for i := 0; i < len(classes); i++ {
        p, t := float64(classes[i][0]), float64(classes[i][1])
        c := class{
            pass: p,
            total: t,
            passRatio: p / t,
            passRatioGain: ((p+1)/(t+1)) - (p/t),
        }

        heap.Push(h, c)
    }

    for extraStudents > 0 {
        c := heap.Pop(h).(class)
        c.pass += 1.0
        c.total += 1.0
        c.passRatio = c.pass / c.total
        c.passRatioGain = (c.pass+1)/(c.total+1) - (c.pass/c.total)
        heap.Push(h, c)
        extraStudents--
    }
    
    sum := 0.0
    n := float64(h.Len())
    for h.Len() > 0 {
        c := heap.Pop(h).(class)
        sum += c.passRatio
    }

    return sum / n
}

// greedy, heap
// time: O((n+m)log(n))
// space: O(n)

type class struct {
    pass, total float64
    passRatio float64
    passRatioGain float64
}

type maxHeap []class

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].passRatioGain > h[j].passRatioGain }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
    *h = append(*h, x.(class))
}
func (h *maxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}