func numberOfPairs(points [][]int) int {
    count := 0
    for i := 0; i < len(points); i++ {
        for j := 0; j < len(points); j++ {
            if i == j { continue }
            
            ax, ay := points[i][0], points[i][1]
            bx, by := points[j][0], points[j][1]
            if ax <= bx && ay >= by {
                valid := true
                for k := 0; k < len(points); k++ {
                    if k == i || k == j { continue }
                    
                    kx, ky := points[k][0], points[k][1]
                    if ax <= kx && kx <= bx && by <= ky && ky <= ay {
                        valid = false
                        break
                    }
                }

                if valid { count++ }
            }
        }
    }

    return count
}

// brute force
// time: O(n^3)
// space: O(1)