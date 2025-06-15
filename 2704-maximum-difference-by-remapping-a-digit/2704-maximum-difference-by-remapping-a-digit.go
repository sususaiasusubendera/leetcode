func minMaxDifference(num int) int {
	digits := intToDigits(num)
    
    // find max num
    // change the first non-nine digit (and all digits that are the same as it) in digits to nine
    maxDigits := make([]int, len(digits))
    copy(maxDigits, digits)
    firstNonNine := -1
    for _, digit := range digits {
        if digit != 9 {
            firstNonNine = digit
            break
        }
    }
    if firstNonNine != -1 {
        for i := 0; i < len(maxDigits); i++ {
            if maxDigits[i] == firstNonNine {
                maxDigits[i] = 9
            }
        }
    }

    // find min num
    // change the first non-zero digit (and all digits that are the same as it) in digits to zero
    minDigits := make([]int, len(digits))
    copy(minDigits, digits)
    firstNonZero := -1
    for _, digit := range digits {
        if digit != 0 {
            firstNonZero = digit
            break
        }
    }
    if firstNonZero != -1 {
        for i := 0; i < len(minDigits); i++ {
            if minDigits[i] == firstNonZero {
                minDigits[i] = 0
            }
        }
    }

    maxNum := digitsToInt(maxDigits)
    minNum := digitsToInt(minDigits)

    return maxNum - minNum
}

// time: O(n)
// space: O(n)

func intToDigits(num int) []int {
    digits := []int{}
    for num > 0 {
        digits = append([]int{num%10}, digits...)
        num /= 10
    }

    return digits
}

func digitsToInt(digits []int) int {
    num := 0
    for _, digit := range digits {
        num = num*10 + digit
    }

    return num
}