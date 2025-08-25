func minimumSum(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    ans := int(1e9)
    var eval func(regionOf func(r, c int) int)
    eval = func(regionOf func(r, c int) int) {
        b0, b1, b2 := newBBox(), newBBox(), newBBox()
        for i := 0; i < rows; i++ {
            for j := 0; j < cols; j++ {
                if grid[i][j] == 1 {
                    id := regionOf(i, j)
                    switch id {
                        case 0:
                            b0.add(i, j)
                        case 1:
                            b1.add(i, j)
                        case 2:
                            b2.add(i, j)
                        default:
                    }
                }
            }
        }
        if b0.count == 0 || b1.count == 0 || b2.count == 0 {
            return
        }
        sum := b0.area() + b1.area() + b2.area()
        if sum < ans {
            ans = sum
        }
    }

    // 3 border vertical
    for c1 := 0; c1 < cols-2; c1++ {
        for c2 := c1+1; c2 < cols-1; c2++ {
            eval(func(r, c int) int {
                if c <= c1 {
                    return 0
                } else if c <= c2 {
                    return 1
                }
                return 2
            })
        }
    }

    // 3 border horizontal
    for r1 := 0; r1 < rows-2; r1++ {
        for r2 := r1+1; r2 < rows-1; r2++ {
            eval(func(r, c int) int {
                if r <= r1 {
                    return 0
                } else if r <= r2 {
                    return 1
                }
                return 2
            })
        }
    }

    // T-shape border
    for rr := 0; rr < rows-1; rr++ {
        for cc := 0; cc < cols-1; cc++ {
            // T-top
            eval(func(r, c int) int {
                if r <= rr && c <= cc {
                    return 0 // top-left
                } else if r <= rr && c > cc {
                    return 1 // top-right
                }
                return 2 // bot
            })

            // T-bot
            eval(func(r, c int) int {
                if r > rr && c > cc {
                    return 0 // bot-right
                } else if r > rr && c <= cc {
                    return 1 // bot-left
                }
                return 2 // top
            })

            // T-right
            eval(func(r, c int) int {
                if r <= rr && c > cc {
                    return 0 // right-top
                } else if r > rr && c > cc {
                    return 1 // right-down
                }
                return 2 // left
            })

            // T-left
            eval(func(r, c int) int {
                if r > rr && c <= cc {
                    return 0 // left-bot
                } else if r <= rr && c <= cc {
                    return 1 // left-top
                }
                return 2 // right
            })
        }
    }

    return ans
}

// array, matrix
// time: O(R^2 + C^2 + RC)
// space: O(RC)

type bBox struct {
    minR, maxR int
    minC, maxC int
    count int
}

func newBBox() bBox {
    return bBox{
        // 1 <= grid.length, grid[i].length <= 30
        minR: 29,
        maxR: 0,
        minC: 29,
        maxC: 0,
    }
}

func (b *bBox) add(r, c int) {
    if r < b.minR {
        b.minR = r
    }
    if r > b.maxR {
        b.maxR = r
    }
    if c < b.minC {
        b.minC = c
    }
    if c > b.maxC {
        b.maxC = c
    }
    b.count++
}

func (b *bBox) area() int {
    if b.count == 0 {
        return 0
    }
    return (b.maxR-b.minR+1) * (b.maxC-b.minC+1)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}