func minOperations(nums []int, k int) int {
    m := map[int]bool{}
    for _, num := range nums {
        if num < k {
            return -1
        }

        if num > k {
            m[num] = true
        }
    }

    return len(m)
}

// hash map
// time: O(n)
// space: O(n)

// p.s. easy tag question, hard tag description