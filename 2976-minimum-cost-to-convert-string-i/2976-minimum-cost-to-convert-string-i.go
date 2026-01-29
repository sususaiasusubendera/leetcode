func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    const INF = 1_000_000_007

	dist := [26][26]int64{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				dist[i][j] = 0
				continue
			}
			dist[i][j] = INF
		}
	}

	for i := 0; i < len(original); i++ {
		u, v, c := int(original[i]-'a'), int(changed[i]-'a'), int64(cost[i])
		dist[u][v] = min(dist[u][v], c)
	}

	// floyd-warshall
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	ans := int64(0)
	for i := 0; i < len(source); i++ {
		s, t := int(source[i]-'a'), int(target[i]-'a')
		if dist[s][t] == INF {
			return -1
		}
		ans += dist[s][t]
	}

	return ans
}

// array, floyd-warshall, graph, string
// time: O(n)
// space: O(1)