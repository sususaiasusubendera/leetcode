func minimumCost(n int, edges [][]int, query [][]int) []int {
    parent := make([]int, n)
    minCost := make([]int, n)
    for i := range parent {
        parent[i] = -1
        minCost[i] = -1
    }

    for _, edge := range edges {
        uRoot, vRoot := find(edge[0], parent), find(edge[1], parent)
        if uRoot != vRoot {
            minCost[uRoot] &= minCost[vRoot]
            parent[vRoot] = uRoot
        }

        minCost[uRoot] &= edge[2]
    }

    result := make([]int, len(query))
    for i, q := range query {
        uRoot, vRoot := find(q[0], parent), find(q[1], parent)
        if q[0] == q[1] {
            result[1] = 0
        } else if uRoot != vRoot {
            result[i] = -1
        } else {
            result[i] = minCost[uRoot]
        }
    }

    return result
}

// NOTICE ME SENPAI!!!
// NOTICE ME SENPAI!!!

func find(node int, parent []int) int {
    if parent[node] < 0 {
        return node
    }

    parent[node] = find(parent[node], parent)
    return parent[node]
}