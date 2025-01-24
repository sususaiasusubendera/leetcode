func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)

    // make a reverse graph
    revGraph := make([][]int, n)
    indegree := make([]int, n) // count incoming edge of every node in revGraph
    for node := 0; node < n; node++ {
        for _, neighbor := range graph[node] {
            revGraph[neighbor] = append(revGraph[neighbor], node)
            indegree[node]++
        }
    }

    // put all nodes with indegree 0 into the queue
    queue := []int{}
    for node := 0; node < n; node++ {
        if indegree[node] == 0 {
            queue = append(queue, node)
        }
    }

    // process the queue with topological sort (Kahn's algorithm)
    safeNodes := []int{}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // dequeue

        safeNodes = append(safeNodes, curr)

        // -1 all the neighbor's indegree in revGraph
        for _, neighbor := range revGraph[curr] {
            indegree[neighbor]--
            // enqueue the neighbor into queue if it's indegree == 0
            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor) 
            }
        }
    }

    sort.Ints(safeNodes)
    return safeNodes
}

// topological sort with queue (indegree) (Kahn's algorithm)
// time: O(n + e)
// space: O(n + e)