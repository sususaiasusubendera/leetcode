func getLongestSubsequence(words []string, groups []int) []string {
    result := []string{}
    for i := 0; i < len(words); i++ {
        if i == 0 || groups[i] != groups[i-1] {
            result = append(result, words[i])
        }
    }

    return result
}

// time: O(n)
// space: O(n)