func removeDuplicates(nums []int) int {
    // constraint: 1 <= nums.length <= 3 * 10^4
    if 1 <= len(nums) && len(nums) <= 30000 {
        if len(nums) == 0 {
            return 0
        }

        u := 1
        for i := 1; i < len(nums); i++ {
            if nums[i] != nums[i-1] {
                nums[u] = nums[i]
                u++
            }
        }
        nums = nums[:u]
        return u
    }
    return -1
} 