func triangularSum(nums []int) int {
    if len(nums) == 1 { return nums[0] }

    n := len(nums)-1
    newNums := make([]int, len(nums)-1)
    for n > 0 {
        for i := 0; i < n; i++ {
            newNums[i] = (nums[i] + nums[i+1]) % 10
        }

        for i := 0; i < n; i++ {
            nums[i] = newNums[i]
        }

        n--
    }

    return newNums[0]
}

// array
// time: O(n^2)
// space: O(n)