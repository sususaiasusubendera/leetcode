func minSubArrayLen(target int, nums []int) int {
    start, total := 0, 0
    minLength := len(nums) + 1
    for end := 0; end < len(nums); end++ {
        total += nums[end]
        for total >= target {
            if end - start + 1 < minLength {
                minLength = end - start + 1
            }
            total -= nums[start]
            start++
        }
    }
    
    if minLength == len(nums) + 1 {
        return 0
    }
    return minLength
}