func minCapability(nums []int, k int) int {
    left, right := 1, int(^uint(0) >> 1)

    for left < right {
        mid := left + (right-left)/2

        count := 0
        prev := -2

        for i := 0; i < len(nums); i++ {
            if nums[i] <= mid && i > prev+1 {
                count++
                prev = i
            }
        }

        if count >= k {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}

// NOTICE ME SENPAI!!!