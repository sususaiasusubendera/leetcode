func stringMatching(words []string) []string {
    seen := map[string]bool{}
    result := []string{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i != j && strings.Contains(words[j], words[i]) {
                if !seen[words[i]] {
                    result = append(result, words[i])
                    seen[words[i]] = true
                }
                break
            }   
        }
    }

    return result
}

// brute force