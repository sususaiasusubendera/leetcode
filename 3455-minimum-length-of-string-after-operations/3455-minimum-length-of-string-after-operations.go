func minimumLength(s string) int {
    m := map[byte]int{}
    length := 0
    for i := 0; i < len(s); i++ {
        m[s[i]]++
        length++
    }

    for _, v := range m {
        if v >= 3 {
            if v % 2 == 0 {
                length -= v - 2
            } else {
                length -= v - 1
            }
        }
    }

    return length
}

// character frequency oriented