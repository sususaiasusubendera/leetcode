func minimumRecolors(blocks string, k int) int {
    countB := 0
    b:= byte('B')

    for i := 0; i < k; i++ {
        if blocks[i] == b {
            countB++
        }
    }

    minOps := k - countB
    if minOps == 0 {
        return minOps
    }

    p1, p2 := 0, k-1
    for p2 < len(blocks)-1 {
        p2++
        if blocks[p2] == b {
            countB++
        }

        if blocks[p1] == b {
            countB--
        }
        p1++

        w := k - countB
        if w < minOps {
            minOps = w
        }
    }

    return minOps
}

// sliding window
// time:O(n)
// space: O(1)