func matchPlayersAndTrainers(players []int, trainers []int) int {
    sort.Ints(players)
    sort.Ints(trainers)

    result := 0
    p := 0
    for i := 0; i < len(players); i++ {
        for p < len(trainers) {
            if players[i] <= trainers[p] {
                result++
                p++
                break
            }
            p++
        }
    }
    
    return result
}

// greedy
// time: O(nlog(n) + mlog(m))
// space: O(1)