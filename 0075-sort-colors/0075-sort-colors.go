func sortColors(nums []int)  {
    freq := make([]int, 3) // freq's size is 3 because 2 is the greatest number in nums; 0 <= nums[i] <= 2
    for _, num := range nums {
        freq[num]++
    }

    idx := 0
    for v, f := range freq {
        for i := 0; i < f; i++ {
            nums[idx] = v
            idx++
        }
    }
}

// counting sort
// time: O(n)
// space: O(1)