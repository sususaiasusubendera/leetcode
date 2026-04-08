func xorAfterQueries(nums []int, queries [][]int) int {
    const MOD = 1_000_000_007

    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]
        for l <= r {
            nums[l] = (nums[l] * v) % MOD
            l += k
        }
    }

    ans := 0
    for _, num := range nums {
        ans ^= num
    }

    return ans
}

// array, brute force
// time: O(qn)
// space: O(1)

// note
// the problem is the solution