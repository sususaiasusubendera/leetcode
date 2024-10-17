func strStr(haystack string, needle string) int {
    // constraint: 1 <= haystack.length, needle.length <= 10^4
    if 1 <= len(haystack) && len(haystack) <= 10000 && 1 <= len(needle) && len(needle) <= 10000 {
        // constraint: haystack and needle consist of only lowercase English characters
        haystack = strings.ToLower(haystack)
        needle = strings.ToLower(needle)

        if len(haystack) == len(needle) {
            if haystack == needle {
                return 0
            }
            return -1
        }

        if len(needle) == 1 {
            for i := 0; i < len(haystack); i++ {
                if haystack[i] == needle[0] {
                    return i
                }
            }
        }

        for i := 0; i <= (len(haystack) - len(needle)); i++ {
            if haystack[i:(i+len(needle))] == needle {
                return i
            }
        }
    }
    return -1
}