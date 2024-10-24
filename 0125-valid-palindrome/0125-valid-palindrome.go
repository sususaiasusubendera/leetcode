func isPalindrome(s string) bool {
    ns := ""
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            ns += strings.ToLower(string(char))
        }
    }

    for i := 0; i < len(ns)/2; i++ {
        if ns[i] != ns[len(ns)-i-1] {
            return false
        }
    }
    return true
}