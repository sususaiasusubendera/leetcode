func containsNearbyDuplicate(nums []int, k int) bool {
    if len(nums) == 1 {
        return false
    }
    
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if lastIdx, exist := m[nums[i]]; exist && i - lastIdx <= k {
            return true
        }
        m[nums[i]] = i
    }

    return false
}