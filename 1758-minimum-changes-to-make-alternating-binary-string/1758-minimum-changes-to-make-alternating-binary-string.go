func minOperations(s string) int {
    // odd: 1, even: 0
    opsZeroEven := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] == '1' {
                opsZeroEven++
            }
        } else {
            if s[i] == '0' {
                opsZeroEven++
            }
        }
    }

    // odd: 0, even: 1
    opsZeroOdd := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] == '0' {
                opsZeroOdd++
            }
        } else {
            if s[i] == '1' {
                opsZeroOdd++
            }
        }
    }

    return min(opsZeroOdd, opsZeroEven)
}

// string
// time: O(n)
// space: O(1)