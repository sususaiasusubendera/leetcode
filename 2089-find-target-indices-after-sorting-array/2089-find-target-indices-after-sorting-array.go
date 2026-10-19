func targetIndices(nums []int, target int) []int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    ans := []int{}
    for i, num := range nums {
        if num == target {
            ans = append(ans, i)
        }
    }

    return ans
}

// array, sorting
// time: O(nlog(n))
// space: O(1)