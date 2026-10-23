func sortPermutation(nums []int) int {
    res := -1
    for i, num := range nums {
        if i != num {
            res &= num
        }
    }

    if res == -1 {
        return 0
    }
    return res
}

// hint 1
// time: O(n)
// space: O(1)