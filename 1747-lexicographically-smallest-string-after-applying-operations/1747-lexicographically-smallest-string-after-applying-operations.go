func findLexSmallestString(s string, a int, b int) string {
    smallest := s
    queue := []string{s}
    visited := map[string]bool{}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // pop
        if visited[curr] { continue }
        visited[curr] = true

        // a
        currA := []byte(curr)
        for i := 0; i < len(currA); i++ {
            if i % 2 == 1 {
                num, _ := strconv.Atoi(string(currA[i]))
                currA[i] = byte((num + a) % 10 + '0')
            }
        }
        stringA := string(currA)
        queue = append(queue, stringA)
        if stringA < smallest { smallest = stringA }

        // b
        currB := []byte(curr)
        lenB := len(currB)
        rotateBy := b % lenB 
        rotated := string(currB[lenB-rotateBy:]) + string(currB[:lenB-rotateBy]) 
        stringB := string(rotated)
        queue = append(queue, stringB)
        if stringB < smallest { smallest = stringB }
    }

    return smallest
}

// bfs, string
// time: O(N * n)
// space: O(N * n)