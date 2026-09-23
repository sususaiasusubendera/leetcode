func minOperations(nums []int, x int) int {
    total := 0
    for _, num := range nums {
        total += num
    }

    target := total - x
    longest := 0
    windowSum := 0
    left := 0

    if target == 0 {
        return len(nums)
    }

    for right := 0; right < len(nums); right++ {
        windowSum += nums[right]

        for windowSum > target && left < len(nums) {
            windowSum -= nums[left]
            left++
        }

        if windowSum == target {
            longest = max(longest, right-left+1)
        }
    }

    if longest == 0 {
        return -1
    }

    return len(nums) - longest
}

// array, sliding window
// time: O(n)
// space: O(1)