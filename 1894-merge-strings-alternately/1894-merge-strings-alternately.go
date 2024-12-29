func mergeAlternately(word1 string, word2 string) string {
    var builder strings.Builder
    p1, p2 := 0, 0 // pointer 1, pointer 2
    l1, l2 := len(word1), len(word2) // length 1, length 2
    for p1 < l1 || p2 < l2 {
        if p1 < l1 {
            builder.WriteByte(word1[p1])
            p1++
        }
        if p2 < l2 {
            builder.WriteByte(word2[p2])
            p2++
        }
    }
    return builder.String()
}