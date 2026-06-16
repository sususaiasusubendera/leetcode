func mapWordWeights(words []string, weights []int) string {
    ans := []byte{}
    for _, word := range words {
        weight := 0
        for i := 0; i < len(word); i++ {
            weight += weights[word[i] - 'a']
        }

        ans = append(ans, byte('z' - (weight % 26)))
    }

    return string(ans)
}

// array, string
// time: O(nk)
// space: O(n)