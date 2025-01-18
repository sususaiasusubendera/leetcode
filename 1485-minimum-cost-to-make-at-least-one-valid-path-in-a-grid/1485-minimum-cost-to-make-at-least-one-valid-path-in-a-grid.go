func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    directions := [][]int{
        {0, 1}, // right
        {0, -1}, // left
        {1, 0}, // down
        {-1, 0}, // up
    }

    // MinHeap BFS-like traversal (priority queue)
    h := &MinHeap{}
    heap.Init(h)
    heap.Push(h, Point{0, 0, 0})

    // visited array
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    for h.Len() > 0 {
        p := heap.Pop(h).(Point)
        x, y, cost := p.x, p.y, p.cost

        // reach bottom-right
        if x == m-1 && y == n-1 {
            return cost
        }

        // skip if already visited
        if visited[x][y] {
            continue
        }
        visited[x][y] = true

        // explore all 4 directions
        for dir, d := range directions {
            nx, ny := x+d[0], y+d[1]
            if nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny] {
                // determine the cost 
                newCost := cost
                if dir+1 != grid[x][y] {
                    newCost++
                }
                heap.Push(h, Point{nx, ny, newCost})
            }
        }
    }

    return -1 // no path is found
}

// TYPES AND METHODS
type Point struct {
    x, y, cost int
}

type MinHeap []Point

// MinHeap methods
func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return h[i].cost < h[j].cost
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Point))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}