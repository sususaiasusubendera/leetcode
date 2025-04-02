func maximumTripletValue(nums []int) int64 {
    result := 0
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            for k := j+1; k < len(nums); k++ {
                temp := ((nums[i] - nums[j]) * nums[k])
                if temp > result {
                    result = temp
                }
            }
        }
    }

    return int64(result)
}

// brute force (hint)
// time: O(n^3)
// space: O(1)