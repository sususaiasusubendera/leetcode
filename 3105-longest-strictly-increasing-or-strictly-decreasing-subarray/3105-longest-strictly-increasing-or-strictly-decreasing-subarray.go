func longestMonotonicSubarray(nums []int) int {
    ans := 1

    // inc
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if i != j {
                if nums[j] <= nums[j-1] {
                    break
                }
                ans = max(ans, j-i+1)
            }
        }
    }

    // dec
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if i != j {
                if nums[j] >= nums[j-1] {
                    break
                }
                ans = max(ans, j-i+1)
            }
        }
    }

    return ans
}

// array
// time: O(n^2)
// space: O(1)