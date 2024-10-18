func removeDuplicates(nums []int) int {
    // constraint: 1 <= nums.length <= 3 * 10^4
    if 1 <= len(nums) && len(nums) <= 30000 {
        if len(nums) == 0 {
            return 0
        }

        j := 1
        for i := 1; i < len(nums); i++ {
            if nums[i] != nums[i-1] {
                nums[j] = nums[i]
                j++
            }
        }
        nums = nums[:j]
        return j
    }
    return -1
} 