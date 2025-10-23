func hasSameDigits(s string) bool {
    intDigits := make([]int, len(s))
    for i := range s {
        intDigits[i] = int(s[i] - '0')
    }

    for len(intDigits) > 2 {
        temp := make([]int, len(intDigits)-1)
        for i := 0; i < len(intDigits)-1; i++ {
            temp[i] = (intDigits[i] + intDigits[i+1]) % 10
        }
        intDigits = temp
    }

    return intDigits[0] == intDigits[1]
}

// math, string
// time: O(n^2)
// space: O(n)