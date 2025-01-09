func countPrefixSuffixPairs(words []string) int {
    count := 0
    for i := 0; i < len(words); i++ {
        for j := i+1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                count++
            }
        }
    }
    return count
}

func isPrefixAndSuffix(sub, str string) bool {
    if len(sub) <= len(str) && strings.HasPrefix(str, sub) && strings.HasSuffix(str, sub) {
        return true
    }
    return false
}

// prefix suffix checker built-in function in Go