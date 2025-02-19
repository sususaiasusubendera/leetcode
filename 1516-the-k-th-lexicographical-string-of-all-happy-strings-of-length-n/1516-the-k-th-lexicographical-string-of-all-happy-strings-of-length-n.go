func getHappyString(n int, k int) string {
    happyStrings := []string{}
    generateHappyStrings(n, "", &happyStrings)

    if len(happyStrings) < k {
        return ""
    }

    return happyStrings[k-1]
}

func generateHappyStrings(n int, currStr string, happyStrings *[]string) {
    // base case
    if len(currStr) == n {
        *happyStrings = append(*happyStrings, currStr)
        return
    }

    // explore
    setChars := []string{"a", "b", "c"}
    for _, char := range setChars {
        if len(currStr) > 0 && string(currStr[len(currStr)-1]) == char {
            continue
        }

        generateHappyStrings(n, currStr+char, happyStrings)
    }
}