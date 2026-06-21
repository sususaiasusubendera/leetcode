func maxIceCream(costs []int, coins int) int {
    iceCream := 0
    totalCost := 0
    sort.Ints(costs)
    for i := 0; i < len(costs); i++ {
        if totalCost + costs[i] > coins {
            return iceCream
        }
        totalCost += costs[i]
        iceCream++
    }
    return iceCream
}

// array, greedy
// time: O(nlog(n))
// space: O(1)