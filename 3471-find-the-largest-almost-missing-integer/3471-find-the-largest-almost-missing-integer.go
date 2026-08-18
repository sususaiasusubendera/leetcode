func largestInteger(nums []int, k int) int {
    freq := make([]int, 51)
    for i := 0; i <= len(nums) - k; i++ {
        seen := map[int]bool{}
        for j := i; j < i + k; j++ {
            if !seen[nums[j]] {
                seen[nums[j]] = true
                freq[nums[j]]++
            }
        }
    }

    ans := -1
    for num, f := range freq {
        if f == 1 {
            ans = max(ans, num)
        }
    }

    return ans
}

// array, hash map
// time: O(nk)
// space: O(k)