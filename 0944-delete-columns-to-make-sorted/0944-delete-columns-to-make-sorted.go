func minDeletionSize(strs []string) int {
    if len(strs) == 1 {
        return 0
    }

    countDelete := 0
    for j := 0; j < len(strs[0]); j++ {
        for i := 1; i < len(strs); i++ {
            if strs[i][j] < strs[i-1][j] {
                countDelete++
                break
            }
        }
    }

    return countDelete
}

// array, string
// time: O(nm)
// space: O(1)