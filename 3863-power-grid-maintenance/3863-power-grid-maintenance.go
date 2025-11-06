func processQueries(c int, connections [][]int, queries [][]int) []int {
	// build an adjency list
	adj := make([][]int, c+1) // 1-indexed
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// find all components using dfs
	componentId := make([]int, c+1)
	compCount := 0
	visited := make([]bool, c+1)

	var dfs func(int, int)
	dfs = func(node int, comp int) {
		visited[node] = true
		componentId[node] = comp
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				dfs(neighbor, comp)
			}
		}
	}

	for i := 1; i <= c; i++ {
		if !visited[i] {
			compCount++
			dfs(i, compCount)
		}
	}

	// prepare heap per component
	heaps := make([]*MinHeap, compCount+1)
	for i := 1; i <= compCount; i++ {
		heaps[i] = &MinHeap{}
	}
	for i := 1; i <= c; i++ {
		heap.Push(heaps[componentId[i]], i)
	}

	// all stations online at the start
	online := make([]bool, c+1)
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	// run the query
	var result []int
	for _, q := range queries {
		t, x := q[0], q[1] // t: type [1, 2]; x: power station id
		// offline
		if t == 2 {
			online[x] = false
		} else if t == 1 {
			if online[x] {
				result = append(result, x)
				continue
			}
			comp := componentId[x]
			h := *heaps[comp]

			// lazy removal
			for h.Len() > 0 && !online[h[0]] {
				heap.Pop(&h)
			}

			if h.Len() == 0 {
				result = append(result, -1)
			} else {
				result = append(result, h[0])
			}
		}
	}

	return result
}

// array, dfs, graph, heap
// time: O((c + n + m)log(c))
// space: O(c + n + m)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}