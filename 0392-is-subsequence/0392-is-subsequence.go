func isSubsequence(s string, t string) bool {
    j := 0
    count := 0
    for i := 0; i < len(s); i++ {
        for j < len(t) {
            if s[i] == t[j] {
                j++
                count++
                break
            }
            j++
        }
    }
    if len(s) == count {
        return true
    } else {
        return false
    }
}