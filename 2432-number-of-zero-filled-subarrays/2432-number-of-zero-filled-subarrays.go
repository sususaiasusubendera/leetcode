func zeroFilledSubarray(nums []int) int64 {
    result := 0
    zeroCounter := 0
    p := 0 // pointer
    for p < len(nums) {
        for p < len(nums) && nums[p] == 0 {
            zeroCounter++
            p++
        }

        result += (zeroCounter * (zeroCounter+1)) / 2
        zeroCounter = 0
        p++
    }

    return int64(result)
}

// math
// time: O(n)
// space: O(1)