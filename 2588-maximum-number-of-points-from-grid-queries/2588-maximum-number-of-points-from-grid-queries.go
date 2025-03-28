func maxPoints(grid [][]int, queries []int) []int {
    m, n, qLen := len(grid), len(grid[0]), len(queries)
    result := make([]int, qLen)
    dirs := [][]int{
        {0, 1},     
        {0, -1},     
        {1, 0},      
        {-1, 0},     
    }

    newQueries := make([][2]int, qLen)
    for i := range queries {
        newQueries[i] = [2]int{queries[i], i}
    }

    sort.Slice(newQueries, func(i, j int) bool {
        return newQueries[i][0] < newQueries[j][0]
    })

    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    curCount := 0
    visited[0][0] = true
    h := &minHeap{}
    heap.Init(h)
    heap.Push(h, cell{0, 0, grid[0][0]})

    for _, q := range newQueries {
        query, idx := q[0], q[1]
        for h.Len() > 0 && (*h)[0].val < query {
            top := heap.Pop(h).(cell)
            curCount++

            for _, d := range dirs {
                newI, newJ := top.i + d[0], top.j + d[1]
                if newI >= 0 && newJ >= 0 && newI < m && newJ < n && !visited[newI][newJ] {
                    heap.Push(h, cell{newI, newJ, grid[newI][newJ]})
                    visited[newI][newJ] = true
                }
            }
        }

        result[idx] = curCount
    }

    return result
}

// NOTICE ME SENPAIII!!!

type Point struct {x, y int}

type cell struct {
    i, j int
    val  int
}

type minHeap []cell

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
    *h = append(*h, x.(cell))
}

func (h *minHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}