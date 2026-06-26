func assignEdgeWeights(edges [][]int) int {
    const MOD = 1_000_000_007

	// if a tree has n nodes, it has n - 1 edges. if a tree has m edges, it has m + 1 nodes
	n := len(edges) + 1

	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	maxDepth := 0

	var dfs func(node, parent, depth int)
	dfs = func(node, parent, depth int) {
		if depth > maxDepth {
			maxDepth = depth
		}

		for _, neigh := range adj[node] {
			if neigh != parent {
				dfs(neigh, node, depth+1)
			}
		}
	}

	dfs(1, 0, 0) // start from node 1

    // ans = 2^(k-1) = number of ways with total cost is odd
    ans := 1
    for i := 0; i < maxDepth-1; i++ {
        ans = (ans * 2) % MOD
    }

    return ans
}

// dfs, math, tree
// time: O(n)
// space: O(n)