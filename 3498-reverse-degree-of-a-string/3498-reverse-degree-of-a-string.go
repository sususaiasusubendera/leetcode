func reverseDegree(s string) int {
    dict := make([]int, 26)
    n := 26
    for i := 0; i < len(dict); i++ {
        dict[i] = n
        n--
    }

    ans := 0
    for i, r := range s {
        ans += dict[int(r - 'a')] * (i + 1)
    }

    return ans
}

// string
// time: O(s)
// space: O(1)