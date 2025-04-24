func countCompleteSubarrays(nums []int) int {
    n := len(nums)
    mapUnique := map[int]bool{}
    for _, num := range nums {
        mapUnique[num] = true
    }
    
    count := 0
    totalUnique := len(mapUnique)
    for left := 0; left < n; left++ {
        mapFreq := map[int]int{}
        for right := left; right < n; right++ {
            mapFreq[nums[right]]++
            if len(mapFreq) == totalUnique {
                count++
            }
        }
    }

    return count
}

// sliding window
// time: O(n^2)
// space: O(n)