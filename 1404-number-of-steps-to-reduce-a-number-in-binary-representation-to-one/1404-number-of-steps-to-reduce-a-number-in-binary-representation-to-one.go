func numSteps(s string) int {
    ops := 0
    carry := 0
    for i := len(s)-1; i > 0; i-- {
        digit := int(s[i] - '0') + carry
        if digit % 2 == 1 {
            ops += 2
            carry = 1
        } else {
            ops++
        }
    }

    return ops + carry
}

// string, bit manipulation, simulation
// time: O(n)
// space: O(1)