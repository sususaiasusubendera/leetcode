func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
    adj := make([][]int, n) // adjacency list
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        adj[a] = append(adj[a], b)
        adj[b] = append(adj[b], a)
    }

    countSplit := 0
    visited := make([]bool, n)
    var dfs func(int) int
    dfs = func(u int) int {
        visited[u] = true
        sum := values[u]

        for _, v := range adj[u] {
            if !visited[v] {
                sum += dfs(v)
            }
        }

        if sum % k == 0 {
            // cut subtree
            countSplit++
            return 0
        }

        return sum
    }

    dfs(0) // root
    return countSplit
}

// tree, dfs
// time: O(n)
// space: O(n)