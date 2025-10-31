func getSneakyNumbers(nums []int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }

    ans := []int{}
    for k, v := range freq {
        if v == 2 {
            ans = append(ans, k)
        }
    }

    return ans
}

// array, hash map
// time: O(n)
// space: O(n)