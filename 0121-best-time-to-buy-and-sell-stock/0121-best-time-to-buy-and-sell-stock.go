func maxProfit(prices []int) int {
    min := prices[0]
    maxp := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        } else {
            p := prices[i] - min
            if p > maxp {
                maxp = p
            }
        }
    }
    return maxp

    // TIME LIMIT EXCEEDED
    // if len(prices) == 1 {
    //     return 0
    // }

    // p := 0
    
    // for i := 0; i < len(prices)-1; i++ {
    //     for j := i+1; j < len(prices); j++ {
    //         if prices[i] > prices[j] {
    //             break
    //         } else {
    //             if prices[j] - prices[i] >= p {
    //                 p = prices[j] - prices[i]
    //             }
    //         }
    //     }
    // }
}