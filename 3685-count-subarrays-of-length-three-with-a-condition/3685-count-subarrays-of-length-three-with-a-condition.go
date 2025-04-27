func countSubarrays(nums []int) int {
    count := 0
    for i := 0; i < len(nums)-2; i++ {
        if float64(nums[i] + nums[i+2]) == float64(nums[i+1])/2 {
            count++
        }
    }

    return count
}

// time: O(n)
// space: O(1)