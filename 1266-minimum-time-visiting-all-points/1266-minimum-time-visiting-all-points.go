func minTimeToVisitAllPoints(points [][]int) int {
    if len(points) == 1 {
        return 0
    }

    ans := 0
    for i := 1; i < len(points); i++ {
        x1, y1 := points[i-1][0], points[i-1][1]
        x2, y2 := points[i][0], points[i][1]

        dx := abs(x2 - x1)
        dy := abs(y2 - y1)

        ans += max(dx, dy)
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// array, geometry, math
// time: O(n)
// space: O(1)