func minimumPushes(word string) int {
    freq := map[byte]int{}
    for i := 0; i < len(word); i++ {
        freq[word[i]]++
    }

    letters := []byte{}
    for letter := range freq {
        letters = append(letters, letter)
    }

    sort.Slice(letters, func(i, j int) bool {
        return freq[letters[i]] > freq[letters[j]]
    })

    keypads := make([]int, 26)
    count := 0
    ans := 0
    for i := 0; i < len(letters); i++ {
        c := letters[i]
        if keypads[c - 'a'] == 0 {
            keypads[c - 'a'] = (count / 8) + 1
            ans += keypads[c - 'a'] * freq[c]
            count++
        } else {
            ans += keypads[c - 'a'] * freq[c]
        }
    }

    return ans
}

// greedy, hash map, sorting, string
// time: O(n + klog(k))
// space: O(k)