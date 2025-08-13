func isPowerOfThree(n int) bool {
    // k = n^3
    maxK := 1162261467 // max k <= (2^31)-1
    return n > 0 && maxK % n == 0
}

// math
// time: O(1)
// space: O(1)