func threeConsecutiveOdds(arr []int) bool {
    countOdd := 0
    for i := 0; i < len(arr); i++ {
        if arr[i] % 2 == 1 {
            countOdd++
        } else {
            countOdd = 0
        }

        if countOdd == 3 {
            return true
        }
    }

    return false
}

// time: O(n)
// space: O(1)