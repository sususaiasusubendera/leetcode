func minOperations(boxes string) []int {
    result := make([]int, len(boxes))

    leftBalls, leftOps := 0, 0
    for i := 0; i < len(boxes); i++ {
        result[i] += leftOps
        if boxes[i] == '1' {
            leftBalls++
        }
        leftOps += leftBalls
    }

    rightBalls, rightOps := 0, 0
    for i := len(boxes)-1; i >= 0; i-- {
        result[i] += rightOps
        if boxes[i] == '1' {
            rightBalls++
        }
        rightOps += rightBalls
    }

    return result
}

// temp