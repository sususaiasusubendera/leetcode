func subsetXORSum(nums []int) int {
    return sumXOR(nums, 0, 0)
}

func sumXOR(nums []int, i, currXOR int) int {
    if i == len(nums) {
        return currXOR
    }

    withElement := sumXOR(nums, i+1, currXOR^nums[i])
    withoutElement := sumXOR(nums, i+1, currXOR)

    return withElement + withoutElement
}

// notice me senpai!!!