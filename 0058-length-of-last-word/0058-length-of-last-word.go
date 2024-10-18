func lengthOfLastWord(s string) int {
    // constraint: 1 <= s.length <= 10^4
    if 1 <= len(s) && len(s) <= 10000 {
        // splitting s string by " "
        ss := strings.Split(s, " ")

        // return the last word length
        for i := len(ss) - 1; i >= 0; i-- {
            if ss[i] != " " && ss[i] != "" {
                return len(ss[i])
            }
        }
    }
    return -1
}