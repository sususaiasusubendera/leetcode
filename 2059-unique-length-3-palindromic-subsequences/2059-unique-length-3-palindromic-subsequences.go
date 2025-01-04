func countPalindromicSubsequence(s string) int {
    order := make([][2]int, 26)   
    for i := 0; i < len(order); i++ {
        order[i] = [2]int{-1, -1}
    }

    for i, c := range s {
        idx := c - 'a'
        if order[idx][0] != -1 {
            order[idx][1] = i
        } else {
            order[idx][0] = i
        }
    }

    count := 0
    for _, pos := range order {
        if pos[1] - pos[0] > 1 {
            unique := make(map[rune]struct{})
            for _, mid := range s[pos[0]+1 : pos[1]] {
                unique[mid] = struct{}{}
            }
            count += len(unique)
        }
    }

    return count
}

// temp