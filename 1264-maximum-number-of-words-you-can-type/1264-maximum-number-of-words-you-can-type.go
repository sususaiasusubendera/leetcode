func canBeTypedWords(text string, brokenLetters string) int {
    brokenLettersMap := map[byte]bool{}
    for i := 0; i < len(brokenLetters); i++ {
        bl := brokenLetters[i]
        brokenLettersMap[bl] = true
    }

    words := 1
    for i := 0; i < len(text); i++ {
        letter := text[i]
        if letter == ' ' { words++ }
    }

    brokenWords := 0
    left := 0
    for right := 0; right <= len(text)-1; right++ {
        if text[right] == ' ' || right == len(text)-1 {
            fullyType := true
            for left <= right {
                letter := text[left]
                if brokenLettersMap[letter] {
                    fullyType = false
                    left = right
                    break
                }
                left++
            }

            if !fullyType { brokenWords++ }
        }
    }

    return words - brokenWords
}

// hash map, string
// time: O(n + b)
// space: O(b)