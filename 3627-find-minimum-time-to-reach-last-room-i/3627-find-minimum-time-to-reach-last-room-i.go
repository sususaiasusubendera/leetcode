func minTimeToReach(moveTime [][]int) int {
    n, m := len(moveTime), len(moveTime[0])
    visited := make([][]bool, n)
    for i := range n {
        visited[i] = make([]bool, m)
    }
    time := 1
    q := [][2]int{[2]int{0, 0}}
    visited[0][0] = true
    for len(q) > 0 {
        var newQ [][2]int
        minWaitTime := int(1e9)
        for _, room := range q {
            l := [2]int{room[0], room[1]-1}
            r := [2]int{room[0], room[1]+1}
            u := [2]int{room[0]-1, room[1]}
            d := [2]int{room[0]+1, room[1]}
            dirs := [][2]int{l,r,u,d}
            needWait := false
            for _, d := range dirs {
                if d[0] < 0 || d[0] > n-1 || d[1] < 0 || d[1] > m-1 || visited[d[0]][d[1]] {
                    continue
                }
                if moveTime[d[0]][d[1]] < time {
                    if d[0] == n-1 && d[1] == m-1 {
                        return time
                    }
                    newQ = append(newQ, d)
                    visited[d[0]][d[1]] = true
                } else {
                    needWait = true
                }
                if moveTime[d[0]][d[1]] < minWaitTime {
                    minWaitTime = moveTime[d[0]][d[1]]
                }
            }
            if needWait {
                newQ = append(newQ, room)
            }
        }
        if minWaitTime > time {
            time = minWaitTime
        }
        q = newQ
        time++
    }
    return -1
}

// NOTICE ME SENPAIII!!!!!