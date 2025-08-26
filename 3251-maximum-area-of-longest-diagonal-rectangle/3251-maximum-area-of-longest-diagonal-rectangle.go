func areaOfMaxDiagonal(dimensions [][]int) int {
    res := -1
    longestDiagonal := -1.0
    for i := 0; i < len(dimensions); i++ {
        l, w := dimensions[i][0], dimensions[i][1]
        diagonal := math.Sqrt(float64((l * l) + (w * w)))
        if diagonal > longestDiagonal {
            longestDiagonal = diagonal
            res = dimensions[i][0] * dimensions[i][1]
        } else if diagonal == longestDiagonal {
            area := l * w
            if area > res {
                res = area
            }
        }
    }

    return res
}

// time: O(n)
// space: O(1)