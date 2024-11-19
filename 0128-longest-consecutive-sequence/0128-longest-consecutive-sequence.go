func longestConsecutive(nums []int) int {
    m := map[int]bool{} // hash set
    for i := 0; i < len(nums); i++ {
        if _, exist := m[nums[i]]; !exist {
            m[nums[i]] = true
        }
    }

    longest := 0
    for key := range m {
        if _, exist := m[key-1]; !exist {
            curr := key
            count := 1
            for m[curr+1] {
            curr++
            count++
            }
            if count > longest {
                longest = count
            }
        }
    }

    return longest
}