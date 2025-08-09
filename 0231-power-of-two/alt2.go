func isPowerOfTwo(n int) bool {
    return 0 < n && n & (n-1) == 0
}

// time: O(1)
// space: O(1)
