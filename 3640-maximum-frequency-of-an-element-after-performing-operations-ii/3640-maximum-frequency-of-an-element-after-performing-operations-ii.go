func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	
    ans := 0
    left, right, idx := 0, 0, 0

    // CASE 1: target is an existing number in nums
    for idx < len(nums) {
        // move 'left' until nums[left] >= nums[idx]-k
        // all elements in [left, idx) can potentially become nums[idx]
        for left < idx && nums[left] < nums[idx]-k {
            left++
        }

        // count numbers equal to nums[idx]
        countSame := 1
        for idx+countSame < len(nums) && nums[idx] == nums[idx+countSame] {
            countSame++
        }
        
        // move 'right' until nums[right] > nums[idx]+k
        // all elements in (idx, right) can potentially become nums[idx]
        for right < len(nums) && nums[right] <= nums[idx]+k {
            right++
        }

        // elements within [left, right) are the window candidates
        // total elements in window  = right - left
        // total elements before idx = idx - left
        // total elements after idx  = right - (idx + countSame)
        // total elements that can be changed = total element before idx + total element after idx
        //                                    = (idx - left) + (right - (idx + countSame))
        //                                    = idx - left + right - idx - countSame
        //                                    = right - left - countSame
        // we can only perform numOperations changes
        // frequency for this target = countSame + min(numOperations, right - left - countSame)
        ans = max(ans, countSame + min(numOperations, right - left - countSame))

        idx += countSame // skip over identical elements
    }

    // CASE 2: target may be a new number (not in nums)
    left = 0
    for right := 0; right < len(nums); right++ {
        // shrink window if nums[right] and nums[left] aren't overlapping
        for left < right && nums[right] - nums[left] > k * 2 {
            left++
        }

        ans = max(ans, min(right - left + 1, numOperations))
    }

    return ans
}

// greedy, sliding window, sorting
// time: O(n log(n))
// space: O(1)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}