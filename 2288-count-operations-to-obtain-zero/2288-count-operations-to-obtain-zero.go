func countOperations(num1 int, num2 int) int {
    countOp := 0 // count operation
    for num1 > 0 && num2 > 0 {
        countOp++
        if num1 >= num2 {
            num1 -= num2
            continue
        }
        num2 -= num1
    }

    return countOp
}

// math, simulation
// time: O(log(min(num1, num2)))
// space: O(1)