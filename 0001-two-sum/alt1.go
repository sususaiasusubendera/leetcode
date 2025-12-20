func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        c := target - nums[i]

        if _, exist := m[c]; exist {
            return []int{m[c], i}
        }

        m[nums[i]] = i
    }
    return nil
}
