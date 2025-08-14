func largestGoodInteger(num string) string {
    left := 0
    digitCount := 1
    maxNum := ""
    for right := 1; right < len(num); right++ {
        if num[right] == num[left] {
            digitCount++
            if digitCount == 3 {
                subNum := num[left:right+1]
                if subNum > maxNum {
                    maxNum = subNum
                }
            }
            continue
        }
        
        left = right
        digitCount = 1
    }

    return maxNum
}

// two pointer
// time: O(n)
// space: O(1)