func binaryGap(n int) int {
    pos := []int{}
    idx := 0
    for n > 0 {
        if n & 1 == 1 {
            pos = append(pos, idx)
        }
        idx++
        n >>= 1
    }

    if len(pos) == 0 {
        return 0
    }

    ans := 0
    for i := 1; i < len(pos); i++ {
        ans = max(ans, pos[i] - pos[i-1])
    }

    return ans
}

// bit manipulation
// time: O(log(n))
// space: O(log(n))