func highestPeak(isWater [][]int) [][]int {
    m, n := len(isWater), len(isWater[0])
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right

    // make matrix 'height', initialize all the values with -1
    h := make([][]int, m) // h: heights
    for i := 0; i < m; i++ {
        h[i] = make([]int, n)
        for j := 0; j < n; j++ {
            h[i][j] = -1
        }
    }

    queue := [][]int{}
    // put all sources (value 1 in the isWater) into queue
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if isWater[i][j] == 1 {
                h[i][j] = 0 // set each water cell to be 0 in matrix 'height'
                queue = append(queue, []int{i, j})
            }
        }
    }

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // dequeue
        for _, dir := range directions {
            newRow, newCol := curr[0]+dir[0], curr[1]+dir[1]
            // check if the new coordinates are valid and unvisited
            if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && h[newRow][newCol] == -1 {
                h[newRow][newCol] = h[curr[0]][curr[1]] + 1
                queue = append(queue, []int{newRow, newCol}) // enqueue
            }
        }
    }

    return h
}

// multi-source bfs
// time: O(mn)
// space: O(mn)

// p.s can directly work on the 'is Water' matrix for a more efficient solution