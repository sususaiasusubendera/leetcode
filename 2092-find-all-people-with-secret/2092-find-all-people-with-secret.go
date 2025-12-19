func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][2] < meetings[j][2]
    })

    parent := make([]int, n)
    know := make([]bool, n)
    for i := range parent {
        parent[i] = i
    }

    know[0], know[firstPerson] = true, true

    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    union := func(a, b int) {
        pa, pb := find(a), find(b)
        if pa != pb {
            parent[pb] = pa
        }
    }

    for i := 0; i < len(meetings); {
        t := meetings[i][2]
        temp := []int{}

        j := i
        for j < len(meetings) && meetings[j][2] == t {
            union(meetings[j][0], meetings[j][1])
            temp = append(temp, meetings[j][0], meetings[j][1])
            j++
        }

        for _, x := range temp {
            if know[x] {
                know[find(x)] = true
            }
        }

        for _, x := range temp {
            know[x] = know[x] || know[find(x)]
        }

        for _, x := range temp {
            parent[x] = x
        }

        i = j
    }

    res := []int{}
    for i := 0; i < n; i++ {
        if know[i] {
            res = append(res, i)
        }
    }
    return res
}

// NOTICE ME SENPAI
// solution from solutions @aashay
// dfs, bfs, union find, graph, sorting