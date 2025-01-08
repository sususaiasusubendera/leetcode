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
    if len(sub) <= len(str) && sub == str[:len(sub)] && sub == str[len(str)-len(sub):] {
        return true
    }
    return false
}