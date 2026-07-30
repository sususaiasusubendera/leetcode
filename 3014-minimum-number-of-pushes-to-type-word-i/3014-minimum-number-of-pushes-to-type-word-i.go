func minimumPushes(word string) int {
    ans := 0
    letters := make([]int, 26)
    count := 0
    for i := 0; i < len(word); i++ {
        c := word[i] - 'a'
        if letters[c] == 0 {
            letters[c] = (count / 8) + 1
            ans += letters[c]
            count++
        } else {
            ans += letters[c]
        }
    }

    return ans
}

// greedy, math, string
// time: O(n)
// space: O(1)