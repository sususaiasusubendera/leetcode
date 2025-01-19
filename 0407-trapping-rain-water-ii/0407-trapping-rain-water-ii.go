func trapRainWater(heightMap [][]int) int {
    m, n := len(heightMap), len(heightMap[0])
    if m < 3 || n < 3 {
        return 0 // no water can be trapped in a map smaller than 3x3
    }

    // visited matrix
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    // MinHeap to store the boundary cells
    h := &MinHeap{}
    heap.Init(h)
    // add all boundary cells to the heap
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || i == m-1 || j == 0 || j == n-1 {
                heap.Push(h, Cell{i, j, heightMap[i][j]})
                visited[i][j] = true
            }
        }
    }

    // directions for exploring neighbors
    directions := []struct{dr, dc int}{
        {-1, 0}, {1, 0}, {0, -1}, {0, 1},
    }

    water := 0
    for h.Len() > 0 {
        cell := heap.Pop(h).(Cell)
        for _, dir := range directions {
            newRow, newCol := cell.row+dir.dr, cell.col+dir.dc
            if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
                visited[newRow][newCol] = true
                // if the current cell is higher than the neighbor, water can be trapped
                water += max(0, cell.height-heightMap[newRow][newCol])
                // push the neighbor into the heap with the maximum height
                heap.Push(h, Cell{newRow, newCol, max(cell.height, heightMap[newRow][newCol])})
            }
        }
    }

    return water
}

// UTILS
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// TYPES AND METHODS
type Cell struct {
    row, col, height int
}

// MinHeap
type MinHeap []Cell

func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return h[i].height < h[j].height
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Cell))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    cell := old[n-1]
    *h = old[:n-1]
    return cell
}