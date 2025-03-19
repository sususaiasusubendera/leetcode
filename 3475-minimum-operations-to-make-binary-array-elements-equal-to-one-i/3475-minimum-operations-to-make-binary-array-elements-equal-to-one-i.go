func minOperations(nums []int) int {
    count := 0
    for i := 2; i < len(nums); i++ {
        if nums[i-2] == 0 {
            count++
            nums[i-2] = 1
            nums[i-1] ^= 1
            nums[i] ^= 1
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

// sliding window, xor
// time: O(n)
// space: O(1)