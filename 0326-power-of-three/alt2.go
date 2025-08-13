func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
    }
    if n == 1 {
        return true
    }
    if n%3 != 0 {
        return false
    }

    return isPowerOfThree(n/3)
}

// time: O(log(n))
// space: O(log(n))
