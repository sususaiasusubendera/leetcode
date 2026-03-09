func checkOnesSegment(s string) bool {
    if len(s) < 2 { return true }

    for i := 1; i < len(s); i++ {
        if s[i-1] == '0' && s[i] == '1' { return false }
    }

    return true
}

// string
// time: O(n)
// space: O(1)