func checkString(s string) bool {
    valid := true
    if s[0] == 'b' {
        for _, c := range s {
            if c == 'a' {
                valid = false
                break 
            }
        }
    } else {
        b := false
        for _, c := range s {
            if !b && c == 'b' {
                b = true
            } else if b && c == 'a' {
                valid = false
                break
            }
        }
    }

    return valid
}

// string
// time: O(n)
// space: O(1)