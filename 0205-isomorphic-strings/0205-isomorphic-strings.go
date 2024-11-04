func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}

    check := true
    for i := 0; i < len(s); i++ {
        if _, exist := m1[s[i]]; !exist {
            m1[s[i]] = t[i]
        } else {
            if m1[s[i]] != t[i] {
                check = false
                break
            }
        }

        if _, exist := m2[t[i]]; !exist {
            m2[t[i]] = s[i]
        } else {
            if m2[t[i]] != s[i] {
                check = false
                break
            }
        }
    }
    return check
}