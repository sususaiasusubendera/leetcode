func spiralOrder(matrix [][]int) []int {
    spiral := []int{}
    if len(matrix) == 1 && len(matrix[0]) == 1 {
        return []int{matrix[0][0]}
    } else if len(matrix) == 1 {
        for j := 0; j < len(matrix[0]); j++ {
            spiral = append(spiral, matrix[0][j])
        }
        return spiral
    } else if len(matrix[0]) == 1 {
        for i := 0; i < len(matrix); i++ {
            spiral = append(spiral, matrix[i][0])
        }
        return spiral

    }

    step := len(matrix) * len(matrix[0])
    xll, yll, xul, yul := 1, 0 ,len(matrix)-1, len(matrix[0])-1
    x, y := 0, 0
    dir := "r"
    for step > 0 {
        if dir == "r" {
            for y <= yul {
                spiral = append(spiral, matrix[x][y])
                y++
                step--
                if y == yul {
                    dir = "d"
                    yul--
                }
            }
        } else if dir == "d" {
            for x <= xul {
                spiral = append(spiral, matrix[x][y])
                x++
                step--
                if x == xul {
                    dir = "l"
                    xul--
                }
            }
        } else if dir == "l" {
            for y >= yll {
                spiral = append(spiral, matrix[x][y])
                y--
                step--
                if y == yll {
                    dir = "u"
                    yll++
                }
            }
        } else if dir == "u" {
            for x >= xll {
                spiral = append(spiral, matrix[x][y])
                x--
                step--
                if x == xll {
                    dir = "r"
                    xll++
                }
            }
        }
    }
    return spiral
}

// xll: x-lower-limit
// yll: y-lower-limit
// xul: x-upper-limit
// yul: y-upper-limit
// dir: direction (r: right, d: down, l: left, u: up)