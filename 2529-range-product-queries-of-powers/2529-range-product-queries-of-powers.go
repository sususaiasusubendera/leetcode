func productQueries(n int, queries [][]int) []int {
    powers := []int{}
    for bit := 0; n > 0; bit++ {
        if n&1 == 1 {
            powers = append(powers, 1<<bit)
        }
        n >>= 1
    }
    
    ans := []int{}
    for _, query := range queries {
        res := 1
        for i := query[0]; i <= query[1]; i++ {
            res = (res * powers[i]) % (1e9+7)
        }

        ans = append(ans, res)
    }

    return ans
}

// bit manipulation
// time: O(nm)
// space: O(n)