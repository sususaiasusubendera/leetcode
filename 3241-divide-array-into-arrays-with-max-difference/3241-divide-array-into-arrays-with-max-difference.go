func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    var ans [][]int

    for i := 0; i < len(nums); i += 3 {
        if nums[i+2]-nums[i] > k {
            return [][]int{}
        }
        ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
    }

    return ans
}

// NOTICE ME SENPAI!!!