func minimumOperations(nums []int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        remainder := nums[i] % 3
        if remainder != 0 { ans++ }
    }

    return ans
}

// array, math
// time: O(n)
// space: O(1)