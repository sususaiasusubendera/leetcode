func maxSum(nums []int) int {
    positiveNumsSet := make(map[int]bool)
    maxNum := nums[0]
    for _, num := range nums {
        if num > 0 {
            positiveNumsSet[num] = true
        }

        if num > maxNum {
            maxNum = num
        }
    }

    if len(positiveNumsSet) == 0 {
        return maxNum
    }

    sum := 0
    for num := range positiveNumsSet {
        sum += num
    }

    return sum
}

// greedy, set
// time: O(n)
// space: O(n)