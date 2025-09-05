func makeTheIntegerZero(num1 int, num2 int) int {
    for k := 1; k < 60; k++ {
        target := num1 - k*num2
        if target < 0 { // no such n^2 < 0
            continue
        }
        // check popcount (total number of 1-bits in bin number) of target <= total operations needed (k)
        // check total operations needed (k) <= target 
        // smallest n^2 is 1, so it's impossible if k > target
        // e.g. k = 3, target = 2. can't reach the target: 1 + 1 + 1 != 2
        if bits.OnesCount64(uint64(target)) <= k && k <= target {
            return k
        }
    }
    return -1
}

// bit manipulation
// time: O(1)
// space: O(1)