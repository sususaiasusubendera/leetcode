func coloredCells(n int) int64 {
    multiplier := 0
    result := 1
    for n != 0 {
        result += (4 * multiplier)
        multiplier++
        n--
    }

    return int64(result)
}

// time: O(n)
// space: O(1)