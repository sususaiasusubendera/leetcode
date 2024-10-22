// alternative (brute force edition) #2

func romanToInt(s string) int {
    n := len(s)
    v := 0
    for i := 0; i < n; i++ {
        if s[i] == 'I' {
            if i+1 < n && s[i+1] == 'V' { // IV
                v += 4
                i++
            } else if i+1 < n && s[i+1] == 'X' { // IX
                v += 9
                i++
            } else { // I
                v += 1
            }
        } else if s[i] == 'X' {
            if i+1 < n && s[i+1] == 'L' { // XL
                v += 40
                i++
            } else if i+1 < n && s[i+1] == 'C' { // XC
                v += 90
                i++
            } else { // X
                v += 10
            }
        } else if s[i] == 'C' {
            if i+1 < n && s[i+1] == 'D' { // CD
                v += 400
                i++
            } else if i+1 < n && s[i+1] == 'M' { // CM
                v += 900
                i++
            } else { // C
                v += 100
            }
        } else if s[i] == 'V' { // V
            v += 5
        } else if s[i] == 'L' { // L
            v += 50
        } else if s[i] == 'D' { // D
            v += 500
        } else if s[i] == 'M' { // M
            v += 1000
        }
    }
    return v
}
