func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    p := make([]int, 26)
    for i := 0; i < 26; i++ {
        p[i] = i
    }

    // build the groups
    for i := 0; i < len(s1); i++ {
        x := int(s1[i]-'a')
        y := int(s2[i]-'a')
        union(x, y, p)
    }

    result := make([]byte, len(baseStr))
    for i := 0; i < len(baseStr); i++ {
        representative := find(int(baseStr[i]-'a'), p)
        result[i] = byte(representative+'a')
    }

    return string(result)
}

// union find
// time: O(n)
// space: O(1)

func find(x int, parent []int) int {
    if parent[x] != x {
        parent[x] = find(parent[x], parent) // path compression
    }

    return parent[x]
}

func union(a, b int, parent []int) {
    pa := find(a, parent)
    pb := find(b, parent)
    if pa == pb {
        return
    }

    // smaller (lexicographically) parent become new one after union
    if pa < pb {
        parent[pb] = pa
    } else {
        parent[pa] = pb
    }
}