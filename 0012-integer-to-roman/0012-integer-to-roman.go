func intToRoman(num int) string {
    s := []string{"I", "V", "X", "L", "C", "D", "M"}

    roman := ""
    c := 0
    for num > 0 {
        d := num % 10
        r := ""
        if d < 4 {
            for i := 0; i < d; i++ {
                r += s[c]
            }
        } else if d == 4 {
            r += s[c] + s[c+1]
        } else if d < 9 {
            r += s[c+1]
            for i := 0; i < d - 5; i++ {
                r += s[c]
            }
        } else {
            r += s[c] + s[c+2]
        }
        roman = r + roman
        num /= 10
        c += 2
    }
    return roman
}