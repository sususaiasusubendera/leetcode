func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)

    ans := int64(0)
    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            w := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
            h := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])

            if w > 0 && h > 0 {
                side := min(w, h)
                area := int64(side) * int64(side)
                ans = max(ans, area)
            }
        }
    }

    return ans
}

// array, geometry, math
// time: O(n^2)
// space: O(1)