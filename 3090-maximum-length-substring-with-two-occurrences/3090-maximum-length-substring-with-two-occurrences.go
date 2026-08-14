func maximumLengthSubstring(s string) int {
    ans := 0
    freq := map[byte]int{}
    left := 0
    for right := 0; right < len(s); right++ {
        freq[s[right]]++

        for freq[s[right]] > 2 {
            freq[s[left]]--
            left++
        }

        ans = max(ans, right - left + 1)
    }

    return ans
}

// hash map, sliding window, string
// time: O(n)
// space: O(k)