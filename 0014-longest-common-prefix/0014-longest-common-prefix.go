func longestCommonPrefix(strs []string) string {
    lcp := strs[0]
    for i := 1; i < len(strs); i++ {
        j := 0
        for j < len(lcp) && j < len(strs[i]) && lcp[j] == strs[i][j] {
            j++
        }
        lcp = lcp[:j]
    }
    return lcp
}