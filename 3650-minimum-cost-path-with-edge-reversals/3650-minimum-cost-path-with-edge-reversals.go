func minCost(n int, edges [][]int) int {
	adj := make([][]Edge, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], Edge{v, w}) // node, weight
		adj[v] = append(adj[v], Edge{u, 2 * w})
	}

	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = 1<<31 - 1
	}

	visited := make([]bool, n)
	h := &MinHeap{&Node{0, 0}}
	for h.Len() > 0 {
		node := heap.Pop(h).(*Node) // current distance, current node
		d, currNode := node.dist, node.val

		if currNode == n-1 {
			return d
		}

		if visited[currNode] {
			continue
		}

		visited[currNode] = true
		for _, neigh := range adj[currNode] {
			neighNode, neighDist := neigh.to, neigh.weight
			newD := d + neighDist
			if newD < dist[neighNode] {
				dist[neighNode] = newD
				heap.Push(h, &Node{
					val:  neighNode,
					dist: newD,
				})
			}
		}
	}

	return -1
}

type Node struct {
	val  int
	dist int
}

type Edge struct {
	to     int
	weight int
}

type MinHeap []*Node

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].dist == h[j].dist {
		return h[i].val < h[j].dist
	}
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(*Node)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// dijkstra, graph, heap
// time: O(nlog(n))
// space: O(n)