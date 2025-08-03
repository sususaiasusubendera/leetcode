func maxTotalFruits(fruits [][]int, startPos int, k int) int {
    left := 0
    sum := 0
    maxFruits := 0
    for right := 0; right < len(fruits); right++ {
        sum += fruits[right][1]

        for left <= right && !isReachable(fruits[left][0], fruits[right][0], startPos, k) {
            sum -= fruits[left][1]
            left++
        }

        if sum > maxFruits {
            maxFruits = sum
        }
    }

    return maxFruits
}

func isReachable(leftPos, rightPos, startPos, k int) bool {
    distLeftToRight := abs(startPos - leftPos) + (rightPos - leftPos)
    distRightToLeft := abs(rightPos - startPos) + (rightPos - leftPos)
    return min(distLeftToRight, distRightToLeft) <= k
}

// sliding window
// time: O(n)
// space: O(1)

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}