func maxFreqSum(s string) int {
    vowel := map[byte]bool{ 'a': true, 'e': true, 'i': true, 'o': true, 'u': true }

    vowelFreq := map[byte]int{}
    consonantFreq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        if vowel[s[i]] {
            vowelFreq[s[i]]++
        } else {
            consonantFreq[s[i]]++
        }
    }

    maxVowel, maxConsonant := 0, 0
    for _, val := range vowelFreq {
        if val > maxVowel { maxVowel = val }
    }

    for _, val := range consonantFreq {
        if val > maxConsonant { maxConsonant = val }
    }

    return maxVowel + maxConsonant
}

// hash map
// time: O(n)
// space: O(1)