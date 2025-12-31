func latestDayToCross(row int, col int, cells [][]int) int {
    n := row * col
    top, bottom := n, n+1
    parent := make([]int, n+2)
    rank := make([]int, n+2)
    land := make([]bool, n)

    for i := range parent { parent[i] = i }

    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    unite := func(a, b int) {
        a, b = find(a), find(b)
        if a == b { return }
        if rank[a] < rank[b] { parent[a] = b
        } else if rank[a] > rank[b] { parent[b] = a
        } else { parent[b] = a; rank[a]++ }
    }

    dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}

    for d := len(cells)-1; d >= 0; d-- {
        r, c := cells[d][0]-1, cells[d][1]-1
        id := r*col + c
        land[id] = true

        for _, di := range dir {
            nr, nc := r+di[0], c+di[1]
            if nr>=0 && nc>=0 && nr<row && nc<col {
                nid := nr*col + nc
                if land[nid] { unite(id, nid) }
            }
        }

        if r == 0 { unite(id, top) }
        if r == row-1 { unite(id, bottom) }

        if find(top) == find(bottom) { return d }
    }
    return 0
}

// NOTICE ME SENPAI!!!
// solution from solutions @aashaypawar
// array, bfs, binary search, dfs, union find