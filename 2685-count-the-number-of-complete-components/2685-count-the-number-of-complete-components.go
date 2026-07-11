func countCompleteComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	for i := 0; i < len(adj); i++ {
		adj[i] = append(adj[i], i)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	m := map[string]int{}
	for i := 0; i < len(adj); i++ {
		sort.Ints(adj[i])
		m[encode(adj[i])]++
	}

	ans := 0
	for k, v := range m {
		if len(decode(k)) == v {
			ans++
		}
	}

	return ans
}

func encode(nums []int) string {
	parts := make([]string, len(nums))
	for i, v := range nums {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

func decode(s string) []int {
	if s == "" {
		return []int{}
	}

	parts := strings.Split(s, ",")
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}

	return nums
}

// adjacency list approach, graph
// time: O(mlog(n))
// space: O(n + m)

// #1 2025/03/22