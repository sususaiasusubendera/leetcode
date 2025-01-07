func minOperations(boxes string) []int {
    result := make([]int, len(boxes))

    // left pass: prefix sum
    leftBalls, leftOps := 0, 0
    for i := 0; i < len(boxes); i++ {
        result[i] += leftOps
        if boxes[i] == '1' {
            leftBalls++
        }
        leftOps += leftBalls
    }

    // right pass: suffix sum
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

// two pass prefix sum and suffix sum