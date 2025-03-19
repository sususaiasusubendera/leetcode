func minOperations(nums []int) int {
    count := 0
    for i := 2; i < len(nums); i++ {
        if nums[i-2] == 0 {
            count++
            nums[i-2] = 1
            ops(&nums[i-1])
            ops(&nums[i])
        }
    }

    sum := 0
    for _, num := range nums {
        sum += num
    }

    if sum == len(nums) {
        return count
    }

    return -1
}

// sliding window
// time: O(n)
// space: O(1)

// helper
func ops(n *int) {
    if *n == 1 {
        *n = 0
    } else {
        *n = 1
    }
}