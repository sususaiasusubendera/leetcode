func minCost(colors string, neededTime []int) int {
    time := 0
    maxTime := neededTime[0]
    for i := 1; i < len(colors); i++ {
        if colors[i] == colors[i-1] {
            time += min(maxTime, neededTime[i])
            maxTime = max(maxTime, neededTime[i])
            continue
        }
        maxTime = neededTime[i]
    }

    return time
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

// array, greedy
// time: O(n)
// space: O(1)