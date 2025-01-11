func canConstruct(s string, k int) bool {
    if len(s) < k {
        return false
    }

    cf := map[rune]int{} // cf: character frequency
    for _, c := range s {
        cf[c]++
    }

    cof := 0 // con: count odd frequency
    for _, f := range cf {
        if f % 2 == 1 {
            cof++
        }
    }

    return cof <= k
}