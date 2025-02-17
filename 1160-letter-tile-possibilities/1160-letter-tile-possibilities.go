func numTilePossibilities(tiles string) int {
    // tile frequencies
    mapFreq := map[byte]int{}
    for i := 0; i < len(tiles); i++ {
        mapFreq[tiles[i]]++
    }

    return backtrack(mapFreq)
} 

// recursive function to explore all the sequences possibilities
func backtrack(m map[byte]int) int {
    total := 0

    for c, freq := range m {
        if freq > 0 {
            m[c]--
            total += 1 + backtrack(m)
            m[c]++
        }
    }

    return total
}

// backtracking approach
// time: O(n!)
// space: O(n)