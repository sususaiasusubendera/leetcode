func lexicalOrder(n int) []int {
    result := make([]int, n)
    curr := 1
    for i := 0; i < n; i++ {
        result[i] = curr
        if curr*10 <= n {
            // move to the first child in lexi order
            curr *= 10
        } else {
            if curr >= n {
                // if curr node has no right sibling:
                // move up to parent
                curr /= 10
            }
            // move to right sibling
            curr++
            for curr%10 == 0 {
                // if no more right siblings at this level:
                // keep going up
                curr /= 10 
            }
        }
    }

    return result
}

// DFS
// time: O(n)
// space: O(n)