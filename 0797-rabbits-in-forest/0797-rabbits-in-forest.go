func numRabbits(answers []int) int {
    freqs := map[int]int{}
    for _, answer := range answers {
        freqs[answer]++
    }

    total := 0
    for x, freq := range freqs {
        groupSize := x + 1
        groupCount := (freq + groupSize - 1) / groupSize
        total += groupCount * groupSize
    }

    return total
}

// time: O(n)
// space: O(n)