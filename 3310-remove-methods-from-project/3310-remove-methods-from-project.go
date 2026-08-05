func remainingMethods(n int, k int, invocations [][]int) []int {
    // make an adjacency list
    adj := make([][]int, n)
    for _, inv := range invocations {
        u, v := inv[0], inv[1]
        adj[u] = append(adj[u], v)
    }

    // find all sus methods
    queue := []int{k}
    sus := make([]bool, n)
    sus[k] = true
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:] // pop
        for _, v := range adj[u] {
            if !sus[v] {
                sus[v] = true
                queue = append(queue, v)
            }
        }
    }

    // find an edge from non-sus to sus
    for _, inv := range invocations {
        u, v := inv[0], inv[1]
        if !sus[u] && sus[v] {
            // can't delete the sus group, return all nodes
            res := make([]int, n)
            for i := 0; i < n; i++ {
                res[i] = i
            }
            
            return res
        }
    }

    // if it's safe, return all non-sus methods
    res := []int{}
    for i := 0; i < n; i++ {
        if !sus[i] {
            res = append(res, i)
        }
    }

    return res
}

// bfs, graph
// time: O(m + n)
// space: O(m + n)