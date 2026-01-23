func minBitwiseArray(nums []int) []int {
    // e.g. num = 23
    // bin: 10111 (23)
    // so,
    // x: 10011 (19)
    // x + 1: 10100 (20)
    // x | (x + 1)
    // 10011
    // 10100 OR
    // 10111 -> 19 OR 20 = 23

    ans := make([]int, len(nums))
    for i, num := range nums {
        // ans[i] = -1 if lsb 0
        if num & 1 == 0 {
            ans[i] = -1
            continue
        }

        // find the last bin digit 1 from lsb
        d := 1 // single-bit mask
        for (num & d) != 0 {
            d <<= 1
        }
        d >>= 1
        ans[i] = num - d
    }

    return ans
}

// time: O(n)
// space: O(n)