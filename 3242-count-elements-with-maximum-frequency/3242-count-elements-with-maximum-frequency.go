func maxFrequencyElements(nums []int) int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }

    max := 0
    for _, v := range freq {
        if v > max { max = v }
    }

    ans := 0
    for _, v := range freq {
        if v == max { ans += max }
    }

    return ans
}

// hash map
// time: O(n)
// space: O(n)