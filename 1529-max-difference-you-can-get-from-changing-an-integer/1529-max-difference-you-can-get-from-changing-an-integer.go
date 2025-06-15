func maxDiff(num int) int {
    digits := intToDigits(num)

    // find a
    // change the first non-nine digit (and all digits that are the same as it) in digits to nine
    maxDigits := make([]int, len(digits))
    copy(maxDigits, digits)
    xMax := -1
    for _, digit := range digits {
        if digit != 9 {
            xMax = digit
            break
        }
    }
    if xMax != -1 {
        for i := 0; i < len(maxDigits); i++ {
            if maxDigits[i] == xMax {
                maxDigits[i] = 9
            }
        }
    }

    // find b
    minDigits := make([]int, len(digits))
    copy(minDigits, digits)
    xMin := -1
    var yMin int
    for i, digit := range digits {
        if i == 0 {
            if digit != 1 {
                xMin = digit
                yMin = 1
                break
            }
        } else {
            if digit != 0 && digit != digits[0] {
                xMin = digit
                yMin = 0
                break
            }
        }
    }
    if xMin != -1 {
        for i := 0; i < len(minDigits); i++ {
            if minDigits[i] == xMin {
                minDigits[i] = yMin
            }
        }
    }

    a := digitsToInt(maxDigits)
    b := digitsToInt(minDigits)

    return a - b
}

// greedy
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