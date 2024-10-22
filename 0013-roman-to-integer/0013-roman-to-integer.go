func romanToInt(s string) int {
    // type Roman struct {
    //     I: 1,
    //     V: 5,
    //     X: 10,
    //     L: 50,
    //     C: 100,
    //     D: 500,
    //     M: 1000
    // }

    n := len(s)
    v := 0
    for i := 0; i < n; i++ {
        if s[i] == 'I' {
            if i+1 < n && s[i+1] == 'V' {
                v += 4
                i++
            } else if i+1 < n && s[i+1] == 'X' {
                v += 9
                i++
            } else {
                v += 1
            }
        } else if s[i] == 'X' {
            if i+1 < n && s[i+1] == 'L' {
                v += 40
                i++
            } else if i+1 < n && s[i+1] == 'C' {
                v += 90
                i++
            } else {
                v += 10
            }
        } else if s[i] == 'C' {
            if i+1 < n && s[i+1] == 'D' {
                v += 400
                i++
            } else if i+1 < n && s[i+1] == 'M' {
                v += 900
                i++
            } else {
                v += 100
            }
        } else if s[i] == 'V' {
            v += 5
        } else if s[i] == 'L' {
            v += 50
        } else if s[i] == 'D' {
            v += 500
        } else if s[i] == 'M' {
            v += 1000
        }
    }
    return v
}