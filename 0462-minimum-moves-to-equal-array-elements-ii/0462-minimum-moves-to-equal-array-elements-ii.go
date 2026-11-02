func minMoves2(nums []int) int {
    med := findMedian(nums)
    res := 0
    for _, num := range nums {
        res += abs(med-num)
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func findMedian(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    if len(nums) % 2 == 1 {
        return nums[len(nums)/2]
    }
    return (nums[len(nums)/2] + nums[(len(nums)/2)-1]) / 2
}

// array, math, median, sorting
// time: O(nlog(n))
// space: O(1)