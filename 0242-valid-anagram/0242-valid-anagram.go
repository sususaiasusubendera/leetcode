func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[rune]int{} // handel unicode characters also (before: map[byte]int{})
    for _, c := range s {
        m[c]++
    }

    for _, c := range t {
        if _, exist := m[c]; exist {
            m[c]--
        }
    }

    for _, count := range m {
        if count != 0 {
            return false
        }
    }

    return true
}
