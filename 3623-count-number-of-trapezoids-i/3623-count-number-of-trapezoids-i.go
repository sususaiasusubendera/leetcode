func countTrapezoids(points [][]int) int {
    const mod = 1_000_000_007
    ans, sum := 0, 0
    pY := map[int]int{} // map y -> point count on y
    
    for _, point := range points {
        y := point[1]
        pY[y]++
    }

    for _, p := range pY {
        c := (p * (p - 1)) / 2 // C(p,2): number of point pairs on this y
        ans = (ans + (c * sum)) % mod
        sum = (sum + c) % mod
    }

    return ans
}

// array, geometry, hash map, math
// time: O(n)
// space: O(m)