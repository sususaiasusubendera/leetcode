func checkStrings(s1 string, s2 string) bool {
    n := len(s1)
    odd1, even1 := []byte{}, []byte{}
    odd2, even2 := []byte{}, []byte{}
    for i := 0; i < n; i++ {
        if i % 2 == 0 {
            even1 = append(even1, s1[i])
            even2 = append(even2, s2[i])
            continue
        }

        odd1 = append(odd1, s1[i])
        odd2 = append(odd2, s2[i])
    }

    sort.Slice(odd1, func(i, j int) bool { return odd1[i] < odd1[j] })
    sort.Slice(even1, func(i, j int) bool { return even1[i] < even1[j] })
    sort.Slice(odd2, func(i, j int) bool { return odd2[i] < odd2[j] })
    sort.Slice(even2, func(i, j int) bool { return even2[i] < even2[j] })

    return string(odd1) == string(odd2) && string(even1) == string(even2)
}

// sorting, string
// time: O(nlog(n))
// space: O(n)