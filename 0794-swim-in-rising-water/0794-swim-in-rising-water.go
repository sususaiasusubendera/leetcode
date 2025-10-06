func swimInWater(grid [][]int) int {
	n := len(grid)

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up, right, down, left

	h := &MinHeap{}
	heap.Push(h, Cell{grid[0][0], 0, 0})
	visited[0][0] = true
	for h.Len() > 0 {
		curr := heap.Pop(h).(Cell)
		if curr.row == n-1 && curr.col == n-1 {
			return curr.cost
		}

		for _, dir := range dirs {
			nextRow, nextCol := curr.row+dir[0], curr.col+dir[1]
			if nextRow >= 0 && nextRow <= n-1 && nextCol >= 0 && nextCol <= n-1 && !visited[nextRow][nextCol] {
				nextCost := max(curr.cost, grid[nextRow][nextCol])
				heap.Push(h, Cell{nextCost, nextRow, nextCol})

				visited[nextRow][nextCol] = true
			}
		}
	}

	return -1
}

// dijkstra, heap, matrix
// time: O(n^2 log(n))
// space: O(n^2)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Cell struct {
	cost int
	row  int
	col  int
}

// min heap implementation
type MinHeap []Cell

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Less(i, j int) bool { return (*h)[i].cost < (*h)[j].cost }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(Cell)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

