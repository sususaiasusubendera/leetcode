func minOperations(s string) int {
    // odd: 0, even: 1
    opsZeroOdd := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] == '1' {
                opsZeroOdd++
            }
        } else {
            if s[i] == '0' {
                opsZeroOdd++
            }
        }
    }

    opsZeroEven := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] == '0' {
                opsZeroEven++
            }
        } else {
            if s[i] == '1' {
                opsZeroEven++
            }
        }
    }

    return min(opsZeroOdd, opsZeroEven)
}

// string
// time: O(n)
// space: O(1)