func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    ans := 2 * n // max possible result if all seats aren't reserved

    sort.Slice(reservedSeats, func(i, j int) bool {
        return reservedSeats[i][0] < reservedSeats[j][0]
    })

    idx := 0
    for idx < len(reservedSeats) {
        currRow := reservedSeats[idx][0]
        // left: s2, s3 reserved
        // midLeft: s4, s5 reserved
        // midRight: s6, s7 reserved
        // right: s8, s9 reserved
        left, midLeft, midRight, right := false, false, false, false
        for idx < len(reservedSeats) && reservedSeats[idx][0] == currRow {
            seat := reservedSeats[idx][1]
            if seat == 2 || seat == 3 {
                left = true
            } else if seat == 4 || seat == 5 {
                midLeft = true
            } else if seat == 6 || seat == 7 {
                midRight = true
            } else if seat == 8 || seat == 9 {
                right = true
            }

            idx++
        }

        if (left && midRight) || (midLeft && midRight) || (midLeft && right) {
            ans -= 2
        } else if left || midLeft || midRight || right {
            ans -= 1
        }
    }

    return ans
}

// array, greedy
// time: O(nlog(n))
// space: O(1)