func maxIncreasingSubarrays(nums []int) int {
	f, g := make([]int, len(nums)), make([]int, len(nums))
	f[0] = 1
	g[len(g)-1] = 1
    
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            g[i] = g[i+1] + 1
        } else {
            g[i] = 1
        }
	}

    ans := 0
    for i := 1; i < len(nums); i++ {
        // f[i-1]: len(nums[a..a+k-1]), g[i]: len(b..b+k-1); b = a + k
        ans = max(ans, min(f[i-1], g[i]))
    }

    return ans
}

// dp (bot-up) two-pass
// time: O(n)
// space: O(n)

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b {  return a }
    return b
}
