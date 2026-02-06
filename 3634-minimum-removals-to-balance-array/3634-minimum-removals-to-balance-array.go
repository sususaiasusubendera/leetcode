func minRemoval(nums []int, k int) int {
    if len(nums) == 1 { return 0 }

    sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
    
    left := 0
    maxWindow := 1
    for right := 1; right < len(nums); right++ {
        for nums[left] * k < nums[right] {
            left++
        }

        if nums[right] <= nums[left] * k {
            maxWindow = max(maxWindow, right - left + 1)
        }
    }

    return len(nums) - maxWindow
}

// array, senior(?), sliding window, sorting
// time: O(nlog(n))
// space: O(1)