func maxSubsequence(nums []int, k int) []int {
    numbers := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        numbers[i] = []int{i, nums[i]}
    }

    sort.Slice(numbers, func(i, j int) bool {
        return numbers[i][1] > numbers[j][1]
    })

    sort.Slice(numbers[:k], func(i, j int) bool {
        return numbers[i][0] < numbers[j][0]
    })

    res := make([]int, k)
    for i := 0; i < k; i++ {
        res[i] = numbers[i][1]
    }

    return res
}

// time: O(nlog(n))
// space: O(n)