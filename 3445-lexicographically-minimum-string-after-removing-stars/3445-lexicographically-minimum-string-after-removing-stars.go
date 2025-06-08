func clearStars(s string) string {
    freq := make([][]int, 26)
    for i := 0; i < len(freq); i++ {
        freq[i] = make([]int, 0)
    }

    sb := []byte(s)
    for i := 0; i < len(sb); i++ {
        if sb[i] != '*' {
            freq[sb[i]-'a'] = append(freq[sb[i]-'a'], i)
        } else {
            for j := 0; j < len(freq); j++ {
                if len(freq[j]) > 0 {
                    last := len(freq[j]) - 1
                    sb[freq[j][last]] = '*'
                    freq[j] = freq[j][:last]
                    break
                }
            }
        }
    }

    result := []byte{}
    for i := 0; i < len(sb); i++ {
        if sb[i] != '*' {
            result = append(result, sb[i])
        }
    }

    return string(result)
}

// greedy
// time: O(n)
// space: O(n)