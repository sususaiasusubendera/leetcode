func rotateTheBox(boxGrid [][]byte) [][]byte {
    m, n := len(boxGrid), len(boxGrid[0])
    for i := 0; i < m; i++ {
        left, right := n-1, n-1
        for right >= 0 && boxGrid[i][right] != '.' {
            right--
            left--
        } 

        for left >= 0 && right >= 0 {
            if boxGrid[i][left] == '#' {
                boxGrid[i][left] = '.'
                boxGrid[i][right] = '#'
                right--
            } else if boxGrid[i][left] == '*' {
                right = left - 1
            }
            left--
        }
    }

    newBox := make([][]byte, n)
    for i := range newBox {
        newBox[i] = make([]byte, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            newBox[i][j] = boxGrid[m-j-1][i]
        }
    }

    return newBox
}

// array, matrix, two pointers
// time: O(mn)
// space: O(nm)