func sortVowels(s string) string {
    vowel := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true,
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }

    temp := []byte{}
    sb := []byte(s)
    for i := 0; i < len(sb); i++ {
        if vowel[sb[i]] {
            temp = append(temp, sb[i])
            sb[i] = '0'
        }
    }

    slices.SortFunc(temp, func(a, b byte) int { return int(a) - int(b) })

    for i := 0; i < len(sb); i++ {
        if sb[i] == '0' {
            sb[i] = temp[0]
            temp = temp[1:]
        }
    }

    return string(sb)
}

// sorting
// time: O(n + mlog(m))
// space: O(n)