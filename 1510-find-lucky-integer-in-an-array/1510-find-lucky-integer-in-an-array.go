func findLucky(arr []int) int {
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {
        freq[arr[i]]++
    }

    result := -1
    for k, v := range freq {
        if k == v && k > result {
            result = k
        }
    }

    return result
}

// time: O(n)
// space: O(n)