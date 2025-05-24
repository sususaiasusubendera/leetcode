func findWordsContaining(words []string, x byte) []int {
    result := []int{}
    for idx, word := range words {
        for i := 0; i < len(word); i++ {
            if word[i] == x {
                result = append(result, idx)
                break           
            }
        }
    }

    return result
}

// time: O(mn)
// space: O(n)