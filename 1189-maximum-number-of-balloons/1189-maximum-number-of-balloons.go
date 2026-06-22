func maxNumberOfBalloons(text string) int {
    freq := map[rune]int{}
    for _, c := range text {
        freq[c]++
    }

    b, a, l, o, n := freq['b'], freq['a'], freq['l'], freq['o'], freq['n']

    count := 0
    for b > 0 && a > 0 && l > 1 && o > 1 && n > 0 {
        count++
        b--
        a--
        l -= 2
        o -= 2
        n--
    }

    return count
}

// hash map, string
// time: O(n)
// space: O(1)