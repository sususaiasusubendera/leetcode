func findRedundantConnection(edges [][]int) []int {
    // initiate nodes with each node being the root of itself
    p := make([]int, len(edges)+1)
    for i := 1; i < len(p); i++ {
        p[i] = i
    }

    for i := 0; ; i++ {
        pa, pb := find(edges[i][0], p), find(edges[i][1], p)
        if pa == pb {
            return edges[i]
        }
        p[pa] = pb
    }
}

// UTIL
func find(x int, p []int) int {
    if p[x] != x {
        p[x] = find(p[x], p)
    }
    return p[x]
}

// disjoint set union (DSU)
// time: O(n)
// space: O(n)

// notice me senpai