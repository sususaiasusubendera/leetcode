func findMissingElements(nums []int) []int {
    sort.Ints(nums)

    ans := []int{}
    num := nums[0]
    idx := 0
    for idx < len(nums) {
        for num != nums[idx] {
            ans = append(ans, num)
            num++
        }
        idx++
        num++
    }

    return ans
}

// array, sorting
// time: O(nlog(n))
// space: O(n)