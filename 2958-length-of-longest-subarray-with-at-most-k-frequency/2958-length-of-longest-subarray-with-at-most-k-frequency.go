func maxSubarrayLength(nums []int, k int) int {
    m := map[int]int{}
    ans := 0
    left := 0
    for right := 0; right < len(nums); right++ {
        m[nums[right]]++

        if m[nums[right]] > k {
            ans = max(ans, right - left)

            for left < right && m[nums[right]] > k {
                m[nums[left]]--
                left++
            }
        }

        if right == len(nums) - 1 {
            ans = max(ans, right - left + 1)
        }
    }

    return ans
}

// array, hash map, sliding window
// time: O(n)
// space: O(n)