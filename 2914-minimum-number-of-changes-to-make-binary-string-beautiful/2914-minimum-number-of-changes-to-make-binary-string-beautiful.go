func minChanges(s string) int {
    count := 0
    for i := 0; i < len(s) / 2; i++ {
        if s[(i*2)] != s[(i*2)+1] {
            count++
        }
    }
    return count
}

// string
// time: O(n/2)
// space: O(1)