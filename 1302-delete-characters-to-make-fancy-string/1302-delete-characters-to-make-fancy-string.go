func makeFancyString(s string) string {
    if len(s) == 1 {
        return s
    }

    p := 0
    count := 1
    result := []byte{}
    for p < len(s)-1 {
        for p < len(s)-1 && s[p] == s[p+1] {
            count++
            p++
        }

        if count >= 2 {
            result = append(result, s[p])
        }
        result = append(result, s[p])

        p++
        if p == len(s)-1 {
            if count < 2 || s[p] != s[p-1] {
                result = append(result, s[p])
            }
        }
        count = 1
    }

    return string(result)
}

// time: O(n)
// space: O(n)