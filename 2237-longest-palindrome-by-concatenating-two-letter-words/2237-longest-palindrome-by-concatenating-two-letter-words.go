func longestPalindrome(words []string) int {
    m := map[string]int{}
    result := 0
    for _, word := range words {
        reversed := string([]byte{word[1], word[0]})
        if freq, exists := m[reversed]; exists && freq > 0 {
            m[reversed]--
            result += 4
        } else {
            m[word]++
        }
    }

    mid := false
    for k, v := range m {
        if v > 0 && k[0] == k[1] {
            mid = true
            break
        }
    }

    if mid {
        return result + 2
    } else {
        return result
    }
}

// time: O(n)
// space: O(n)