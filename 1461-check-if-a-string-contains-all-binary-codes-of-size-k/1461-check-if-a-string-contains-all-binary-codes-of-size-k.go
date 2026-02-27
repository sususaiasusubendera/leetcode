func hasAllCodes(s string, k int) bool {
    h := map[string]bool{}
    for i := 0; i < len(s) - k + 1; i++ {
        if !h[s[i:i+k]] {
            h[s[i:i+k]] = true
        }
    }

    binSubTot := 1
    for i := 0; i < k; i++ {
        binSubTot *= 2
    }

    return len(h) == binSubTot
}

// bit manipulation, hash map, string
// time: O(nk)
// space: O(2^k)