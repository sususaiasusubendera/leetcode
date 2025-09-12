func doesAliceWin(s string) bool {
    vowel := map[byte]bool{ 'a': true, 'e': true, 'i': true, 'o': true, 'u': true }

    vowelCount := 0
    for i := 0; i < len(s); i++ {
        if vowel[s[i]] { vowelCount++ }
    }

    if vowelCount == 0 {
        return false // bob win
    } else if vowelCount % 2 == 1 {
        return true // alice win
    }

    // alice win when vowelCount % 2 == 0
    // 1. alice takes a substring that contains vowelCount-1 vowels
    // 2. bob takes a substring that contains 0 vowels
    // 3. alice takes the rest
    return true
}

// string
// time: O(n)
// space: O(1)