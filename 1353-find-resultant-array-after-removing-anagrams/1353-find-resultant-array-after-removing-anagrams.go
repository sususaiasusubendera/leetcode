func removeAnagrams(words []string) []string {
    idx := 1
    for idx < len(words) {
        if isAnagram(words[idx], words[idx-1]) {
            words = append(words[:idx], words[idx+1:]...)
        } else {
            idx++
        }
    }

    return words
}

// array, sorting, string
// time: O(n * k log(k))
// space: O(k)

func isAnagram(x, y string) bool {
    sliceX, sliceY := []byte(x), []byte(y)
    slices.SortFunc(sliceX, func(a, b byte) int {
        if a > b { return -1 }
        if a < b { return 1 }
        return 0
    })
    slices.SortFunc(sliceY, func(a, b byte) int {
        if a > b { return -1 }
        if a < b { return 1 }
        return 0
    })

    return string(sliceX) == string(sliceY)
}