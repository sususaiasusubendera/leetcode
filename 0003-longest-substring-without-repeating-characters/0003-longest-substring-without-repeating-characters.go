func lengthOfLongestSubstring(s string) int {
    m := map[byte]bool{}
    start, ls := 0, 0
    for end := 0; end < len(s); end++ {
        for m[s[end]] {
            delete(m, s[start])
            start++
        }
        m[s[end]] = true
        if end - start + 1 > ls {
            ls = end - start + 1
        }
    }
    return ls
}