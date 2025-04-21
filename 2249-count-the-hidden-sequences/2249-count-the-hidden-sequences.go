func numberOfArrays(differences []int, lower int, upper int) int {
    var x, y, curr int
    for _, diff := range differences {
        curr += diff
        x, y = min(x, curr), max(y, curr)
        if y-x > upper-lower {
            return 0
        }
    }

    return (upper - lower) - (y - x) + 1
}

// notice me senpai

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}