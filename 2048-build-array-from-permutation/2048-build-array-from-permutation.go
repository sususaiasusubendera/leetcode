func buildArray(nums []int) []int {
    sol := make([]int, len(nums))
    for idx, num := range nums {
        sol[idx] = nums[num]
    }
    return sol
}