func divideArray(nums []int) bool {
    numFreq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        numFreq[nums[i]]++
    }

    for _, value := range numFreq {
        if value % 2 == 1 {
            return false
        }
    }

    return true
}

// time: O(n)
// space: O(n)