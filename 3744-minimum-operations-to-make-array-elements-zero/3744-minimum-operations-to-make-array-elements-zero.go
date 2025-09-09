func minOperations(queries [][]int) int64 {
    var f func(n int) int
    f = func(n int) int {
        if n == 0 { return 0 }

        total := 0
        step := 1 // steps needed for current block
        start := 1 // start of the block
        // each block's length: 3 * 4^(k-1)
        length := 3 // length of the first block (1..3)
        for start <= n {
            end := start + length - 1
            if end > n { end = n } // cut, if the block > n

            count := end - start + 1 // total numbers in the block
            total += count * step // steps contribution from the block

            // next block
            step++
            start *= 4
            length *= 4
        }

        return total
    }

    ops := 0
    for _, query := range queries {
        l, r := query[0], query[1]
        steps := f(r) - f(l-1) // total steps [l..r]
        ops += (steps + 1) / 2 // 1 op can process 2 numbers
    }

    return int64(ops)
}

// array, math
// time: O(qlog(n))
// space: O(1)