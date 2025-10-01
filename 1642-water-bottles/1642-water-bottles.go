func numWaterBottles(numBottles int, numExchange int) int {
    res := numBottles
    for numBottles > 0 {
        remainder := numBottles % numExchange
        numBottles /= numExchange
        res += numBottles

        if numBottles + remainder >= numExchange { numBottles += remainder }
    }

    return res
}

// math
// time: O(log(n))
// space: O(1)