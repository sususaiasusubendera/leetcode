func numberOfSubstrings(s string) int {
    freq := map[byte]int{}
    ans := 0
    left := 0
    for right := 0; right < len(s); right++ {
        freq[s[right]]++
        for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
            ans += len(s) - right
            freq[s[left]]--
            left++
        }
    }

    return ans
}

// hash map, sliding window, string
// time: O(n)
// space: O(1)

// #1 2025/03/11