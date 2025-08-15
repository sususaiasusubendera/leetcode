func isPowerOfFour(n int) bool {
    if n <= 0 {
        return false
    }
    if n == 1 {
        return true
    }
    if n%4 != 0 {
        return false
    }

    return isPowerOfFour(n/4)
}

// time: O(log(n))
// space: O(log(n))
