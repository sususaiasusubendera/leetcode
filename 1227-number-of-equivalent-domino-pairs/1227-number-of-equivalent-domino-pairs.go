func numEquivDominoPairs(dominoes [][]int) int {
    count := 0
    m := map[[2]int]int{}
    for _, domino := range dominoes {
        a, b := domino[0], domino[1]
        var k [2]int
        if a < b {
            k[0], k[1] = a, b
        } else {
            k[0], k[1] = b, a
        }

        count += m[k]
        m[k]++
    }

    return count
}

// hashmap
// time: O(n)
// space: O(k)