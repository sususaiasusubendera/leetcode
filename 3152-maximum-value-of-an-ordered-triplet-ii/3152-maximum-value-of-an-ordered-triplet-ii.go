func maximumTripletValue(nums []int) int64 {
    prefixMax, suffixMax := make([]int, len(nums)), make([]int, len(nums))
    prefixMax[0] = nums[0]
    suffixMax[len(nums)-1] = nums[len(nums)-1]

    // preprocess prefixMax
    for i := 1; i < len(nums); i++ {
        if nums[i] > prefixMax[i-1] {
            prefixMax[i] = nums[i]
            continue
        }
        prefixMax[i] = prefixMax[i-1]
    }

    // preprocess suffixMax
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] > suffixMax[i+1] {
            suffixMax[i] = nums[i]
            continue
        }
        suffixMax[i] = suffixMax[i+1]
    }

    result := 0
    for j := 1; j < len(nums)-1; j++ {
        temp := (prefixMax[j-1] - nums[j]) * suffixMax[j+1]
        if temp > result {
            result = temp
        }
    }

    return int64(result)
}

// prefix suffix
// time: O(n)
// space: O(n)