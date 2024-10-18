func lengthOfLastWord(s string) int {
    // constraint: 1 <= s.length <= 10^4
    if 1 <= len(s) && len(s) <= 10000 {
        i := len(s) - 1
        lastWordLength := 0

        // ignore spaces at the end of the string
        for i >= 0 && s[i] == ' ' {
            i--
        }

        // count the characters of the last word of the string
        for i >= 0 && s[i] != ' ' {
            lastWordLength++
            i--
        }

        return lastWordLength
    }
    return -1
}