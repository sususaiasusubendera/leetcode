func smallestNumber(n int) int {
    bitLen := bits.Len(uint(n))
    return (1 << bitLen) - 1
}

// bit manipulation
// time: O(1)
// space: O(1)