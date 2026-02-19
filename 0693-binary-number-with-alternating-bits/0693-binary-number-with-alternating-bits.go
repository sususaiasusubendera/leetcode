func hasAlternatingBits(n int) bool {
    bit := n & 1
    altBits := true
    for n > 0 {
        n >>= 1
        if bit == n & 1 {
            altBits = false
            break
        }
        bit = n & 1
    }
    return altBits
}

// bit manipulation
// time: O(log(n))
// space: O(1)