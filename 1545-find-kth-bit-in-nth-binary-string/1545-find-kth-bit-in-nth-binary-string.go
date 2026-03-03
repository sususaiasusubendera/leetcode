func findKthBit(n int, k int) byte {
    s := "0"
    for i := 1; i < n; i++ {
        temp := "1" + reverse(invert(s))
        s += temp
    }

    return s[k-1]
}

// simulation, string
// time: O(2^n)
// space: O(2^n)

func reverse(s string) string {
    res := []byte(s)
    for i := 0; i < len(s)/2; i++ {
        res[i], res[len(s)-1-i] = res[len(s)-1-i], res[i]
    }
    return string(res)
}

func invert(s string) string {
    res := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            res[i] = '0'
        } else {
            res[i] = '1'
        }
    }
    return string(res)
}