func spellchecker(wordlist []string, queries []string) []string {
    wordExact := map[string]string{}
    wordLower := map[string]string{}
    wordVowel := map[string]string{}
    for _, word := range wordlist {
        wordExact[word] = word

        lower := strings.ToLower(word)
        if _, exist := wordLower[lower]; !exist { wordLower[lower] = word }

        mask := vowelMask(word)
        if _, exist := wordVowel[mask]; !exist { wordVowel[mask] = word }
    }

    ans := []string{}
    for _, query := range queries {
        if word, exist := wordExact[query]; exist {
            ans = append(ans, word)
        } else if word, exist := wordLower[strings.ToLower(query)]; exist {
            ans = append(ans, word)
        } else if word, exist := wordVowel[vowelMask(query)]; exist {
            ans = append(ans, word)
        } else {
            ans = append(ans, "")
        }
    }

    return ans
}

// array, hash map, string
// time: O(nm + q)
// space: O(nm + q)

func vowelMask(s string) string {
    vowel := map[rune]bool{ 'a': true, 'e': true, 'i': true, 'o': true, 'u': true }
    var b strings.Builder
    for _, c := range s {
        lower := unicode.ToLower(c)
        if vowel[lower] {
            b.WriteByte('*')
        } else {
            b.WriteRune(lower)
        }
    }
    return b.String()
}