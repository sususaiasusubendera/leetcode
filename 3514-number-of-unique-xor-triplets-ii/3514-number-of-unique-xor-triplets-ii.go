func uniqueXorTriplets(nums []int) int {
    // xor operation result can't exceed the longest bit
    // 1500 = 10111011100 -> max: 11111111111 (11) = 2047
    const MAX = 2047

    a, b, c := make([]bool, MAX+1), make([]bool, MAX+1), make([]bool, MAX+1)

    for _, n := range nums {
        a[n] = true
    }

    for i := 0; i < MAX; i++ {
        if !a[i] {
            continue
        }
        for _, n := range nums {
            b[i^n] = true
        }
    }

    for i := 0; i < MAX; i++ {
        if !b[i] {
            continue
        }
        for _, n := range nums {
            c[i^n] = true
        }
    }

    ans := 0
    for _, ok := range c {
        if ok {
            ans++
        }
    }

    return ans
}

// array, brute force + optimization
// time: O(MAX * n)
// space: O(MAX)