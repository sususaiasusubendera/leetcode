func majorityElement(nums []int) int {
    n := len(nums)
    // constraint: 1 <= n <= 5 * 10^4
    if 1 <= n && n <= 50000 {
        count := 0
        var x int
        for i, p := 0, &x; i < n; i++ {
            if count == 0 {
                *p = nums[i]
            }
            
            if nums[i] == *p {
                count++
            } else {
                count --
            }
        }
        return x
    }
    return -1
}