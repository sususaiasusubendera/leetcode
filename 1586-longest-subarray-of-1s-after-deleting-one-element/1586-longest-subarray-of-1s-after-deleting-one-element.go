func longestSubarray(nums []int) int {
    countZero := 0
    longestWindow := 0
    left := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            countZero++
        } 

        for countZero > 1 {
            if nums[left] == 0 {
                countZero--
            }
            left++
        }

        longestWindow = max(longestWindow, right-left)
    }

    return longestWindow
}

// sliding window
// time: O(n)
// space: O(1)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}