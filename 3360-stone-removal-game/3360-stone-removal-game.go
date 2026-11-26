func canAliceWin(n int) bool {
    count := 0
    remove := 10
    for n >= remove {
        n -= remove
        remove--
        count++
    }
    if count % 2 == 0 {
        return false
    }
    return true
}

// math
// time: O(sqrt(n))
// space: O(1)