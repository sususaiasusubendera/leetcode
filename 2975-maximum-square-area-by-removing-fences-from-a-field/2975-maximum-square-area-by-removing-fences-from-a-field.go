func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    // add unremovable fence (boundary) for hFences and vFences
    hFences = append([]int{1}, append(hFences, m)...)
    vFences = append([]int{1}, append(vFences, n)...)

    sort.Slice(hFences, func(i, j int) bool { return hFences[i] < hFences[j] })
    sort.Slice(vFences, func(i, j int) bool { return vFences[i] < vFences[j] })

    hSet := map[int]struct{}{}
    for i := 0; i < len(hFences)-1; i++ {
        for j := i + 1; j < len(hFences); j++ {
            s := hFences[j] - hFences[i]
            if _, exists := hSet[s]; !exists {
                hSet[s] = struct{}{}
            }
        }
    }

    maxS := -1 // max side
    for i := 0; i < len(vFences)-1; i++ {
        for j := i + 1; j < len(vFences); j++ {
            s := vFences[j] - vFences[i]
            if _, exists := hSet[s]; exists {
                maxS = max(maxS, s)
            }
        }
    }

    if maxS == -1 {
        return -1
    }

    const MOD = 1_000_000_007

    return (maxS * maxS) % MOD
}

// array, hash map
// time: O(h^2 + v^2)
// space: O(h^2 + h + v)