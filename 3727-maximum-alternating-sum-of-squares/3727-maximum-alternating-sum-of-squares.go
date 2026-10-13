func maxAlternatingSum(nums []int) int64 {
    sort.Slice(nums, func(i, j int) bool {
        return abs(nums[i]) > abs(nums[j])
    })

    var limit int
    if len(nums) % 2 == 1 {
        limit = (len(nums) / 2) + 1
    } else {
        limit = len(nums) / 2
    }

    res := int64(0)
    for i := 0; i < len(nums); i++ {
        if i < limit {
            res += int64(nums[i] * nums[i])
        } else {
            res -= int64(nums[i] * nums[i])
        }
    }

    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// array, sorting
// time: O(nlog(n))
// space: O(1)