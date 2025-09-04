func findClosest(x int, y int, z int) int {
    p1, p2 := abs(z - x), abs(z - y)
    if p1 < p2 { return 1 }
    if p1 > p2 { return 2 }
    return 0
}

// time: O(1)
// space: O(1)

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}