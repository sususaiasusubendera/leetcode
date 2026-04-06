func robotSim(commands []int, obstacles [][]int) int {
    o := map[[2]int]bool{}
    for _, obstacle := range obstacles {
        o[[2]int{obstacle[0], obstacle[1]}] = true
    }

    ans := 0
    face := 0 // face: 0 (north), 1 (east), 2 (south), 3 (west)
    p := [2]int{0, 0} // p (x, y)
    for _, command := range commands {
        if 1 <= command && command <= 9 {
            switch face {
                // north
                case 0:
                    for i := 1; i <= command; i++ {
                        if !o[[2]int{p[0], p[1]+1}] {
                            p[1]++
                            continue
                        }
                        break
                    }
                // east
                case 1:
                    for i := 1; i <= command; i++ {
                        if !o[[2]int{p[0]+1, p[1]}] {
                            p[0]++
                            continue
                        }
                        break
                    }
                // south
                case 2:
                    for i := 1; i <= command; i++ {
                        if !o[[2]int{p[0], p[1]-1}] {
                            p[1]--
                            continue
                        }
                        break
                    }
                // west
                case 3:
                    for i := 1; i <= command; i++ {
                        if !o[[2]int{p[0]-1, p[1]}] {
                            p[0]--
                            continue
                        }
                        break
                    }
            }
            ans = max(ans, (p[0] * p[0]) + (p[1] * p[1]))
        } else if command == -1 || command == -2 {
            face = turn(command, face)
        }
    }

    return ans
}

func turn(t, f int) int {
    if t == -1 {
        f = (f + 1) % 4
    } else if t == -2 {
        f = (f - 1 + 4) % 4
    }
    return  f
}

// array, hash map
// time: O(n)
// space: O(m)