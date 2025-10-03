func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])

    // no water can be trapped in a heightMap smaller than 3x3
	if m < 3 || n < 3 {
		return 0
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// add all outer layer cells of the heightMap to heap, mark visited
	h := &MinHeap{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, Cell{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up, right, down, left
	water := 0
	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		for _, dir := range dirs {
			newRow, newCol := cell.row+dir[0], cell.col+dir[1]
			if newRow > 0 && newRow < m-1 && newCol > 0 && newCol < n-1 && !visited[newRow][newCol] {
				neighborHeight := heightMap[newRow][newCol]
				if neighborHeight < cell.height {
					water += cell.height - neighborHeight
					neighborHeight = cell.height
				}

				heap.Push(h, Cell{newRow, newCol, neighborHeight})
				visited[newRow][newCol] = true
			}
		}
	}

	return water
}

// array, bfs, heap, matrix
// time: O(mn log(mn))
// space: mn

// min heap implementation
type Cell struct {
	row    int
	col    int
	height int
}

type MinHeap []Cell

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Less(i, j int) bool { return (*h)[i].height < (*h)[j].height }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(Cell)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}