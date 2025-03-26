func minOperations(grid [][]int, x int) int {
    nums := []int{}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            nums = append(nums, grid[i][j])
        }
    }

    slices.Sort(nums)
    target := nums[len(nums)/2]
    count := 0
    for i := range nums {
        if target % x != nums[i] % x {
            return -1
        }

        diff := target - nums[i]
        if diff < 0 {
            diff = -diff
        }

        count += diff / x
    }

    return count
}

// notice me senpai!