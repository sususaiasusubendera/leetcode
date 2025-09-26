func triangleNumber(nums []int) int {
    if len(nums) < 3 { return 0 }

    sort.Ints(nums)
    res := 0
    for k := len(nums)-1; k >= 2; k-- {
        i, j := 0, k-1
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                res += j - i
                j--
            } else {
                i++
            }
        }
    }

    return res
}

// array, two pointers
// time: O(n^2)
// space: O(1)