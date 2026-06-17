func processStr(s string, k int64) byte {
    // find the result's length
	l := int64(0)
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            if l > 0 {
                l--
            }
        } else if s[i] == '#' {
            l *= 2
        } else if s[i] == '%' {
            continue
        } else {
            l++
        }
    }

    if k > l - 1 {
        return '.'
    }

    // process the s backward, reverse-engineer the k
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '*' {
            l++
        } else if s[i] == '#' {
            l /= 2
            if k >= l { 
                k = k - l
            }
        } else if s[i] == '%' {
            k = l - k - 1
        } else {
            if k == l - 1 {
                return s[i]
            } else {
                l--
            }
        }
    }

    return '.'
}

// simulation, string
// time: O(n)
// space: O(1)