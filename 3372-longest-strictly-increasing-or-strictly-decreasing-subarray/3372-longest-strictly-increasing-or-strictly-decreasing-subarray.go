func longestMonotonicSubarray(nums []int) int {
    incLen, decLen := 1, 1
    maxLen := 1
    for i := 1; i < len(nums); i++ {
        // strictly increasing subarray 
        if nums[i] > nums[i-1] {
            incLen++
        } else {
            incLen = 1
        }

        // strictly decreasing subarray
        if nums[i] < nums[i-1] {
            decLen++
        } else {
            decLen = 1
        }

        // find longest subarray by comparing incLen and decLen
        if incLen > decLen {
            if incLen > maxLen {
                maxLen = incLen
            }
        } else {
            if decLen > maxLen {
                maxLen = decLen
            }
        }
    }

    return maxLen
}

// time: O(n)
// space: O(1)