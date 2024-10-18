func majorityElement(nums []int) int {
    n := len(nums)
    // constraint: 1 <= n <= 5 * 10^4
    if 1 <= n && n <= 50000 {
        count := 0
        var x int
        for i := 0; i < n; i++ {
            if count == 0 {
                x = nums[i]
            }
            
            if nums[i] == x {
                count++
            } else {
                count --
            }
        }
        return x
    }
    return -1
}