func maxIncreasingSubarrays(nums []int) int {
	ans := 0
    prevCount := 0
    count := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            count++
        } else {
            prevCount = count
            count = 1
        }

        ans = max(ans, min(prevCount, count)) // adjacent 2 increasing subarrays
        ans = max(ans, count/2) // 1 increasing subarray, split it into 2
    }

    return ans
}

// greedy
// time: O(n)
// space: O(1)

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b {  return a }
    return b
}