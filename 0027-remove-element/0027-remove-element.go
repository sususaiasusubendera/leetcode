func removeElement(nums []int, val int) int {
    // constraint: 0 <= nums.length <= 100
    // constraint: 0 <= val <= 100
    if 0 <= len(nums) && len(nums) <= 100 && 0 <= val && val <= 100 {
        i := 0
        count := 0
        for _, v := range nums {
            if v != val {
                nums[i] = v
                i++
            } else {
                count++
            }
        }
        nums = nums[:i]
        k := i
        return k
    }
    return -1
}