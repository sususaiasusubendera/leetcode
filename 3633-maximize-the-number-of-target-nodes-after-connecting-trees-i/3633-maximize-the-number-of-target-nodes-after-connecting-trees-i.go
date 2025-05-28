func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	var dfs func(node, parent int, children [][]int, k int) int
	dfs = func(node, parent int, children [][]int, k int) int {
		if k < 0 {
			return 0
		}
		res := 1
		for _, child := range children[node] {
			if child == parent {
				continue
			}
			res += dfs(child, node, children, k-1)
		}
		return res
	}

	build := func(edges [][]int, k int) []int {
		n := len(edges) + 1
		children := make([][]int, n)
		for _, edge := range edges {
			u, v := edge[0], edge[1]
			children[u] = append(children[u], v)
			children[v] = append(children[v], u)
		}
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = dfs(i, -1, children, k)
		}
		return res
	}

	n := len(edges1) + 1
	count1 := build(edges1, k)
	count2 := build(edges2, k-1)
	maxCount2 := 0
	for _, c := range count2 {
		if c > maxCount2 {
			maxCount2 = c
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = count1[i] + maxCount2
	}
	return res
}

// editorial
// notice me senpai!!!