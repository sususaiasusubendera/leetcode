func maxProfit(prices []int) int {
    min := prices[0]
    maxp := 0

    for i := 1; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        } else {
            p := prices[i] - min
            maxp += p
            min = prices[i]
        }
    }

    return maxp
}