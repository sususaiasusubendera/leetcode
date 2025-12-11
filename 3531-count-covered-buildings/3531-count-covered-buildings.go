func countCoveredBuildings(n int, buildings [][]int) int {
    mx, my := map[int][]int{}, map[int][]int{}

    for _, p := range buildings {
        x, y := p[0], p[1]
        mx[x] = append(mx[x], y)
        my[y] = append(my[y], x)

        sort.Ints(mx[x])
        sort.Ints(my[y])
    }

    ans := 0
    for _, p := range buildings {
        x, y := p[0], p[1]
        
        if mx[x][0] < y && y < mx[x][len(mx[x])-1] && my[y][0] < x && x < my[y][len(my[y])-1] {
            ans++
        }
    }

    return ans
}

// array, hash map, sorting
// time: O(nlog(n))
// space: O(n)