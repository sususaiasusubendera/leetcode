func repeatedNTimes(nums []int) int {
    target := len(nums) / 2
    freq := map[int]int{}
    for _, num := range nums {
        freq[num]++
        if freq[num] == target {
            return num
        }
    }

    return -1
}

// array, hash map
// time: O(n)
// space: O(n)