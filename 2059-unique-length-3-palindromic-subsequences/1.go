func countPalindromicSubsequence(s string) int {
    // initialize an array with [-1, -1] (for position tracking [start, end])
    order := make([][2]int, 26)   
    for i := 0; i < len(order); i++ {
        order[i] = [2]int{-1, -1}
    }

    // traverse the s, locate 'start index' and 'end index' of s[i] if any
    for i, c := range s {
        idx := c - 'a'
        if order[idx][0] != -1 {
            order[idx][1] = i
        } else {
            order[idx][0] = i
        }
    }

    // count unique middle characters between start and end
    count := 0
    for _, pos := range order {
        // check if any middle characters
        if pos[1] - pos[0] > 1 { 
            // hash set using struct{} is better (space) than bool, CMIIW
            unique := make(map[rune]struct{}) 
            for _, mid := range s[pos[0]+1 : pos[1]] {
                // mark the unique character using struct{}{}
                unique[mid] = struct{}{}
            }
            count += len(unique)
        }
    }

    return count
}
