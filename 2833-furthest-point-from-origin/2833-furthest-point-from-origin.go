func furthestDistanceFromOrigin(moves string) int {
    r, l, x := 0, 0, 0
    for _, move := range moves {
        if move == 'R' {
            r++
        } else if move == 'L' {
            l++
        } else {
            x++
        }
    }

    return max(r, l) - min(r, l) + x
}

// counting, string
// time: O(n)
// space: O(1)