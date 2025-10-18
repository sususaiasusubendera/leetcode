func maxDistinctElements(nums []int, k int) int {
    sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
    ans := 0
    curr := -1 << 63 // the last digit used
    for _, num := range nums {
        low, high := num - k, num + k
        next := max(curr+1, low)
        if next <= high {
            ans++
            curr = next
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b { return a }

    return b
}

// array, greedy, sorting
// time: O(n log(n))
// space: O(1)