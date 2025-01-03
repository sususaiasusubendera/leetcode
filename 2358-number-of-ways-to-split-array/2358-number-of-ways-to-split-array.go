func waysToSplitArray(nums []int) int {
    sumNums := 0
    for i := 0; i < len(nums); i++ {
        sumNums += nums[i]
    }

    count, leftSum := 0, 0
    for i := 0; i < len(nums)-1; i++ {
        leftSum += nums[i]
        if leftSum >= sumNums - leftSum {
            count++
        }
    }

    return count
}