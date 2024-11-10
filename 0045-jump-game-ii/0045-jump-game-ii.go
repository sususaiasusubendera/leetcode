func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }

    end, jump, farthest := 0, 0, 0
    for i := 0; i < len(nums); i++ {
        farthest = max(farthest, i + nums[i])

        if i == end {
            jump++
            end = farthest
        }

        if end >= len(nums)-1 {
            break
        }
    }
    return jump
}

func max(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}