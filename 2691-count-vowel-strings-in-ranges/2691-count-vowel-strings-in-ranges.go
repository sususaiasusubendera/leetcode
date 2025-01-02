func vowelStrings(words []string, queries [][]int) []int {
    m := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    // precompute prefix sum
    prefix := make([]int, len(words))
    for i, word := range words {
        if m[word[0]] && m[word[len(word)-1]] {
            prefix[i] = 1
        }

        if i > 0 {
            prefix[i] += prefix[i-1]
        }
    }

    // process each query
    result := make([]int, len(queries))
    for i, query := range queries {
        li, ri := query[0], query[1]
        if li == 0 {
            result[i] = prefix[ri]
        } else {
            result[i] = prefix[ri] - prefix[li-1]
        }
    }

    return result
}

// time: O(w + q)