func magnificentSets(n int, edges [][]int) int {
    // adjacency list
    adjList := make([][]int, n)
    for _, e := range edges {
        u, v := e[0]-1, e[1]-1 // convert to zero-based index
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }

    maxHeight := make([]int, n)
    for i := range maxHeight {
        maxHeight[i] = -1
    }

    sumDiameters := 0

    for start := 0; start < n; start++ {
        isBipartite, diameter := bfsCheckAndFindHeight(start, adjList, maxHeight)
        if !isBipartite {
            return -1
        }
        if diameter != -1 {
            sumDiameters += diameter
        }
    }

    return sumDiameters
}

// notice me senpai

// UTILS
// bfsCheckAndFindHeight
func bfsCheckAndFindHeight(start int, adjList [][]int, maxHeight []int) (bool, int) {
    level := make([]int, len(adjList))
    queue := []int{start}
    level[start] = 1
    maxLevel := -1
    allHeightsAssigned := true

    for front := 0; front < len(queue); front++ {
        v := queue[front]
        maxLevel = max(maxLevel, maxHeight[v])

        if v != start && maxHeight[v] == -1 {
            allHeightsAssigned = false
        }

        for _, u := range adjList[v] {
            if level[u] == 0 {
                level[u] = level[v] + 1
                queue = append(queue, u)
            } else if absDiff(level[u], level[v]) != 1 {
                return false, -1
            }
        }
    }

    maxHeight[start] = level[queue[len(queue)-1]]

    if allHeightsAssigned {
        maxLevel = max(maxLevel, maxHeight[start])
        return true, maxLevel
    }
    return true, -1
}

// absDiff
func absDiff(a, b int) int {
    if a >= b {
        return a - b
    }
    return b - a
}