func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    r1x1, r1y1, r1x2, r1y2 := rec1[0], rec1[1], rec1[2], rec1[3]
    r2x1, r2y1, r2x2, r2y2 := rec2[0], rec2[1], rec2[2], rec2[3]

    // r1 is left of r2, r1 is right of r2,
    // r1 is below r2, r1 is above r2
    if r1x2 <= r2x1 || r1x1 >= r2x2 || r1y2 <= r2y1 || r1y1 >= r2y2 {
        return false
    }

    return true
}

// geometry, math
// time: O(1)
// space: O(1)