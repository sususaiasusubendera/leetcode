func increasingTriplet(nums []int) bool {
    first, second := math.MaxInt, math.MaxInt
    for i := 0; i < len(nums); i++ {
        if nums[i] <= first {
            first = nums[i]
        } else if nums[i] <= second {
            second = nums[i]
        } else {
            return true
        }
    }
    return false
}

// space: O(1); time: O(n)
// P.S. I found a little logical fallacy (for the index condition) in this solution (for example, this test case: [20, 100, 10, 12, 5, 13]), but OK (CMIIW).
