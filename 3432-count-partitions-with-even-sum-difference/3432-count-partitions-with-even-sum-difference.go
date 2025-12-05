func countPartitions(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    ans := 0
    leftSum, rightSum := 0, sum
    for i := 0; i < len(nums)-1; i++ {
        leftSum += nums[i]
        rightSum -= nums[i]
        res := rightSum - leftSum
        if res % 2 == 0 {
            ans++
        }
    }

    return ans
}

// array
// time: O(n)
// space: O(1)