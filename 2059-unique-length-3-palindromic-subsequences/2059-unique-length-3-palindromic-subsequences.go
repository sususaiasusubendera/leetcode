func countPalindromicSubsequence(s string) int {
    letters := make([][2]int, 26) // store start idx and end idx of a letter
    for i := range letters {
        letters[i] = [2]int{-1, -1} // [2]int{ start idx, end idx }
    }

    for i := 0; i < len(s); i++ {
        c := s[i] - 'a'
        if letters[c][0] == -1 { letters[c][0] = i; continue }
        letters[c][1] = i
    }

    ans := 0
    for _, letter := range letters {
        start, end := letter[0], letter[1]
        if start != -1 && end != -1 && end - start > 1 {
            seen := map[byte]struct{}{}
            for i := start+1; i <= end-1; i++ {
                if _, exists := seen[s[i]]; !exists {
                    seen[s[i]] = struct{}{}
                    ans++
                }
            }
        }
    }

    return ans
}

// hash map, string
// time: O(n)
// space: O(1)