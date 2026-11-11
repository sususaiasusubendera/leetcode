func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
    if need1 == 0 && need2 == 0 {
        return 0
    }

    a, b, c, d := 0, 0, 0, 0
    
    // costBoth (type 1 and type 2) all in
    c = max(need1, need2) * costBoth

    // type 1 all in
    a = need1 * cost1

    // type 2 all in
    b = need2 * cost2

    // costBoth + type 1 or type 2
    minOne := true
    if need1 > need2 {
        minOne = false
    }
    d = min(need1, need2) * costBoth
    if minOne {
        d += (need2-need1) * cost2
    } else {
        d += (need1-need2) * cost1
    }

    return int64(min(min(a+b, c), d))
}

// greedy, math
// time: O(1)
// sapce: O(1)