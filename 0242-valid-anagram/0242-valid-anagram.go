func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[byte]int{}
    for i, _ := range s {
        m[s[i]] += 1
    }

    for i, _ := range t {
        if _, exist := m[t[i]]; exist {
            m[t[i]] -= 1
        }
    }

    for i, _ := range s {
        if m[s[i]] != 0 {
            return false
        }
    }

    return true
}
