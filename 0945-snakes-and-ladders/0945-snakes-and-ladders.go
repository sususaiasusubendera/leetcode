func snakesAndLadders(board [][]int) int {
    n := len(board)

    // translate label (1-indexed) to position (matrix coordinate (row, col))
    labelToPos := func(label int) (int, int) {
        row := n - 1 - ((label-1)/n)
        col := (label-1) % n
        if (n-row) % 2 == 0 {
            col = n - 1 - col
        }

        return row, col
    }

    visited := make([]bool, (n*n)+1)
    queue := []int{1} // start from label 1
    moves := 0

    // bfs
    for len(queue) > 0 {
        nextQueue := []int{}
        for _, curr := range queue {
            if curr == (n*n) {
                return moves
            }

            for dice := 1; dice <= 6 && curr + dice <= (n*n); dice++ {
                next := curr + dice
                row, col := labelToPos(next)
                if board[row][col] != -1 {
                    // meet snake or ladder
                    next = board[row][col]
                }

                if !visited[next] {
                    visited[next] = true
                    nextQueue = append(nextQueue, next)
                }
            }
        }

        queue = nextQueue
        moves++
    }

    return -1
}

// bfs
// time: O(n^2)
// space: O(n^2)