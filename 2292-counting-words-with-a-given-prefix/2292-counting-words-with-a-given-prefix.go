func prefixCount(words []string, pref string) int {
    count := 0
    for i := 0; i < len(words); i++ {
        if len(pref) <= len(words[i]) && pref == words[i][:len(pref)] {
            count++
        }
    }
    return count
}