func maxAscendingSum(nums []int) int {
    maxSum := nums[0]
    sum := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            sum += nums[i]
        } else {
            sum = nums[i]
        }

        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}

// one pass
// time: O(n)
// space: O(1)