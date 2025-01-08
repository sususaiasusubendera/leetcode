func stringMatching(words []string) []string {
    seen := map[string]bool{}
    result := []string{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i != j && kmsContains(words[j], words[i]) {
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

// pattern (substring) matching using KMP algorithm
func kmsContains(text, pattern string) bool {
    if len(pattern) > len(text) {
        return false
    }

    lps := buildLPS(pattern)
    i, j := 0, 0
    for i < len(text) {
        if text[i] == pattern[j] {
            i++
            j++
        }
        if j == len(pattern) { 
            return true // pattern found
        } else if i < len(text) && text[i] != pattern[j] {
            if j != 0 { 
                j = lps[j-1] // 'i' doesn't need to backtrack, use 'j' instead via LPS
            } else {
                i++
            }
        }
    }

    return false
}

// LPS (longest prefix sufix) for word
func buildLPS(s string) []int {
    i := 1
    length := 0
    lps := make([]int, len(s))
    for i < len(s) {
        if s[i] == s[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                i++
            }
        }
    }
    return lps
}

// brute force with style
// KMP (Knuth-Morris-Pratt) algorithm (string matching)