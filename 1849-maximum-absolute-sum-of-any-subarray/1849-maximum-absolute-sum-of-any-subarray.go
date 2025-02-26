func maxAbsoluteSum(nums []int) int {
    currPosSum, currNegSum := 0, 0
    maxAbsSum := 0

    for _, num := range nums {
        currPosSum = max(0, currPosSum + num)
        currNegSum = max(0, currNegSum - num)
        maxAbsSum = max(maxAbsSum, max(currPosSum, currNegSum))
    }

    return maxAbsSum
}

// NOTICE ME SENPAI!!!