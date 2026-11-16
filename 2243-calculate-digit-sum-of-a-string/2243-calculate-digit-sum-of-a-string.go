func digitSum(s string, k int) string {
    for len(s) > k {
        temp := ""
        for i := 0; i < len(s); i += k {
            val := 0
            for j := i; j < len(s) && j < i+k; j++ {
                val += int(s[j]-'0')
            }
            temp += strconv.Itoa(val)
        }
        s = temp
    }
    return s
}

// simulation, string
// time: O(n)
// space: O(n)