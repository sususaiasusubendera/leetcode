func findDiagonalOrder(mat [][]int) []int {
    rows, cols := len(mat), len(mat[0])
    dirUpRight := true // false: dirDownLeft
    p := newPoint()
    for len(p.res) < (rows * cols) {
        p.addRes(mat[p.x][p.y])
        if dirUpRight {
            if p.y == cols-1 { // hit right side
                p.moveDown()
                dirUpRight = false
            } else if p.x == 0 { // hit up side
                p.moveRight()
                dirUpRight = false
            } else {
                p.moveUpRight()
            }
        } else {
            if p.x == rows-1 { // hit down side
                p.moveRight()
                dirUpRight = true
            } else if p.y == 0 { // hit left side
                p.moveDown()
                dirUpRight = true
            } else {
                p.moveDownLeft()
            }
        }
    }

    return p.res
}

// array, matrix
// time: O(mn)
// space: O(mn)

type point struct {
    x int
    y int
    res []int
} 

func newPoint() point {
    return point{x: 0, y: 0}
}

func (p *point) addRes(n int) {
    p.res = append(p.res, n)
}

func (p *point) moveUpRight() {
    p.x--
    p.y++
}

func (p *point) moveRight() {
    p.y++
}

func (p *point) moveDownLeft() {
    p.x++
    p.y--
}

func (p *point) moveDown() {
    p.x++
}