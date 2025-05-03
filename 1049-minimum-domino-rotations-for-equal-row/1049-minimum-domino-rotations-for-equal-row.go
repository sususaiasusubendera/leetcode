func minDominoRotations(tops []int, bottoms []int) int {
    var T, B map[int][2]int
    T = make(map[int][2]int)
    B = make(map[int][2]int)
    n := len(tops)
    ans := math.MaxFloat64
    for i := 0; i < n; i++ {
        T[tops[i]] = [2]int{T[tops[i]][0] + 1, T[tops[i]][1]}
        B[bottoms[i]] = [2]int{B[bottoms[i]][0] + 1, B[bottoms[i]][1]}
        if bottoms[i] != tops[i] {
            T[bottoms[i]] = [2]int{T[bottoms[i]][0], T[bottoms[i]][1] + 1}
            B[tops[i]] = [2]int{B[tops[i]][0], B[tops[i]][1] + 1}
        }
    }
    v := []int{tops[0], tops[n-1], bottoms[0], bottoms[n-1]}
    for _, i := range(v) {
        if T[i][0] + T[i][1] == n {
            ans = math.Min(ans, float64(T[i][1]))
        }
        if B[i][0] + B[i][1] == n {
            ans = math.Min(ans, float64(B[i][1]))
        }
    }
    if ans == math.MaxFloat64 {
        return -1
    }
    return int(ans)
}

// NOTICE ME SENPAI!!!