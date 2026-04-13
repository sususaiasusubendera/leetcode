func getMinDistance(nums []int, target int, start int) int {
    ans := 1_000_000_007
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            ans = min(ans, abs(i - start))
        }
    }

    return ans
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// array
// time: O(n)
// space: O(1)