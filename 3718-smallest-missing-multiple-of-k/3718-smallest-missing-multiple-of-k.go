func missingMultiple(nums []int, k int) int {
    m := map[int]bool{}
    for _, num := range nums {
        m[num] = true
    }

    num := k
    for num < 1000 {
        if !m[num] {
            return num
        }

        num += k
    }

    return num
}

// array, hash map
// time: O(n + t)
// space: O(n)