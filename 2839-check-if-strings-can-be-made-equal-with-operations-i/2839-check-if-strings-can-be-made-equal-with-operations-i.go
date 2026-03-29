func canBeEqual(s1 string, s2 string) bool {
    a1, b1 := []byte{s1[0], s1[2]}, []byte{s1[1], s1[3]}
    a2, b2 := []byte{s2[0], s2[2]}, []byte{s2[1], s2[3]}
    
    // sort for group idx 0, 2 and 1, 3
    sort.Slice(a1, func(i, j int) bool { return a1[i] < a1[j] })
    sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
    sort.Slice(a2, func(i, j int) bool { return a2[i] < a2[j] })
    sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

    return string(a1) == string(a2) && string(b1) == string(b2)
}

// string
// time: O(1)
// space: O(1)