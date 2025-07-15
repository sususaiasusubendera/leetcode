func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }

    hasVow, hasCon := false, false
    for _, char := range word {
        if unicode.IsLetter(char) {
            c := unicode.ToLower(char)
            if vowels[c] {
                hasVow = true
            } else {
                hasCon = true
            }
        } else if !unicode.IsDigit(char) {
            return false
        }
    }

    return hasVow && hasCon
}

// time: O(n)
// space: O(1)

var vowels = map[rune]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}