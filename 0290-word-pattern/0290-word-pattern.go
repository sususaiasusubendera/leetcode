func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)
    if len(pattern) != len(words) {
        return false
    }

    ps := map[byte]string{}
    sp := map[string]byte{}
    for i := 0; i < len(pattern); i++ {
        if _, exist := ps[pattern[i]]; !exist {
            ps[pattern[i]] = words[i]
        }

        if _, exist := sp[words[i]]; !exist {
            sp[words[i]] = pattern[i]
        }
    }

    for i := 0; i < len(pattern); i++ {
        if pattern[i] != sp[words[i]] || words[i] != ps[pattern[i]] {
            return false
        }
    }

    return true
}