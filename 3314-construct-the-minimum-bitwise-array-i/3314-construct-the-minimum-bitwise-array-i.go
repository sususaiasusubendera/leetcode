func minBitwiseArray(nums []int) []int {
    ans := make([]int, len(nums))
    for i, num := range nums {
        val := -1
        for j := 1; j < num; j++ {
            if (j | (j + 1)) == num {
                val = j
                break
            }
        }
        ans[i] = val
    }

    return ans
}

// array, bit manipulation
// time: O(nm)
// space: O(n)