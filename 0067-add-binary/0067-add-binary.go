func addBinary(a string, b string) string {
    idxa, idxb := len(a)-1, len(b)-1
    ans := []byte{}
    carry := false
    for idxa >= 0 || idxb >= 0 {
        if idxa >= 0 && idxb >= 0 && a[idxa] == '1' && b[idxb] == '1' {
            if carry {
                ans = append(ans, '1')
            } else {
                ans = append(ans, '0')
            }
            carry = true
        } else if idxa >= 0 && idxb >= 0 && a[idxa] == '0' && b[idxb] == '0' {
            if carry {
                ans = append(ans, '1')
            } else {
                ans = append(ans, '0')
            }
            carry = false
        } else if idxa >= 0 && idxb >= 0 && ((a[idxa] == '1' && b[idxb] == '0') || (a[idxa] == '0' && b[idxb] == '1')) {
            if carry {
                ans = append(ans, '0')
                carry = true
            } else {
                ans = append(ans, '1')
                carry = false
            }
        } else if idxa >= 0 {
            if a[idxa] == '1' {
                if carry {
                    ans = append(ans, '0')
                    carry = true
                } else {
                    ans = append(ans, '1')
                    carry = false
                }
            } else {
                if carry {
                    ans = append(ans, '1')
                } else {
                    ans = append(ans, '0')
                }
                carry = false
            }
        } else if idxb >= 0 {
            if b[idxb] == '1' {
                if carry {
                    ans = append(ans, '0')
                    carry = true
                } else {
                    ans = append(ans, '1')
                    carry = false
                }
            } else {
                if carry {
                    ans = append(ans, '1')
                } else {
                    ans = append(ans, '0')
                }
                carry = false
            }
        }

        idxa--
        idxb--
    }

    if carry {
        ans = append(ans, '1')
    }

    reverse(ans)

    return string(ans)
}

func reverse(b []byte) {
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
}

// bit manipulation, math, simulation, string
// time: O(n)
// space: O(n)