func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}

    for i := 0; i < len(s); i++ {
        if v, exist := m1[s[i]]; exist && v != t[i] {
            return false
        }
        m1[s[i]] = t[i]

        if v, exist := m2[t[i]]; exist && v != s[i] {
            return false
        }
        m2[t[i]] = s[i]
    }
    return true
}
