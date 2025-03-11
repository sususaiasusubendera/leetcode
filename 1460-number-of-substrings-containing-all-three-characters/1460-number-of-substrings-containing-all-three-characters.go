func numberOfSubstrings(s string) int {
    freq := map[byte]int{
        'a': 0,
        'b': 0,
        'c': 0,
    }

    left := 0
    result := 0
    for right := 0; right < len(s); right++ {
        freq[s[right]]++

        for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
            result += len(s) - right
            freq[s[left]]--
            left++
        }
    }

    return result
}

// time: O(n)
// space: O(1)