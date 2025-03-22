func countCompleteComponents(n int, edges [][]int) int {
	edgeMap := map[int][]int{}
	res := 0
	for _, e := range edges {
		edgeMap[e[0]] = append(edgeMap[e[0]], e[1])
		edgeMap[e[1]] = append(edgeMap[e[1]], e[0])
	}
	var dfs func(int)
	visited := make([]bool, n)
	ec := 0
	vc := 0
	dfs = func(v int) {
		ec += len(edgeMap[v])
        vc++ 
        visited[v] = true
        for  _, toV := range edgeMap[v] {
            if !visited[toV] {
                dfs(toV)
            }
        }
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			ec = 0
			vc = 0
			dfs(i)
			if ec == vc*(vc-1) {
				res++
			}
		}
	}
	return res
}

// NOTICE ME SENPAIII!!!