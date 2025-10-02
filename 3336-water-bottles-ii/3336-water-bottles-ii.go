func maxBottlesDrunk(numBottles int, numExchange int) int {
    drink := numBottles
    fullBottles := 0
    emptyBottles := numBottles
    for fullBottles > 0 || emptyBottles >= numExchange {
        if emptyBottles >= numExchange {
            fullBottles++
            emptyBottles -= numExchange
            numExchange++
            continue
        } else {
            drink += fullBottles
            emptyBottles += fullBottles
            fullBottles = 0
        }
    }

    return drink
}

// math, simulation
// time: O(n)
// space: O(1)