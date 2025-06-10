func maxDifference(s string) int {
    freq := make([]int, 26)
    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
    }

    oddFreq, evenFreq := -1, -1
    for i := 0; i < 26; i++ {
        if freq[i] % 2 == 1 && freq[i] != 0 {
            if oddFreq == -1 {
                oddFreq = freq[i]
                continue
            }

            if freq[i] > oddFreq {
                oddFreq = freq[i]
            }
        } else if freq[i] % 2 == 0 && freq[i] != 0 {
            if evenFreq == -1 {
                evenFreq = freq[i]
                continue
            }

            if freq[i] < evenFreq {
                evenFreq = freq[i]
            }
        }
    }

    return oddFreq - evenFreq
}

// time: O(n)
// space: O(1)