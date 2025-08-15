func isPowerOfFour(n int) bool {
    // 0x55555555 = 01010101 01010101 01010101 01010101
    // check if it's power of two
    // the power of 4 always has a bit 1 in an odd position (LSB: 0th, 2th, 4th, ...)
    return n > 0 && (n & (n-1)) == 0 && (n & 0x55555555) != 0
}

// time: O(1)
// space: O(1)