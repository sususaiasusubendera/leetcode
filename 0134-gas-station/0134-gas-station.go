func canCompleteCircuit(gas []int, cost []int) int {
    start, tank, checkFuel := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        checkFuel += gas[i] - cost[i]
        tank += gas[i] - cost[i]
        if tank < 0 {
            start = i + 1
            tank = 0
        }
    }

    if checkFuel < 0 {
        return -1
    }

    return start
}