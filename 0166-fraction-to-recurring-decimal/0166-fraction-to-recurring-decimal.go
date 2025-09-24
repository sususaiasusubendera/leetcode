func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 { return "0" }

    ans := ""

    // handle the sign
    if (numerator < 0) != (denominator < 0) { ans += "-" }

    // convert to positive number if negative (to make the operation easier)
    if numerator < 0 { numerator = -numerator }
    if denominator < 0 { denominator = -denominator }

    // integer part
    integer := numerator / denominator
    ans += strconv.Itoa(integer)

    // decimal part
    remainder := numerator % denominator
    if remainder == 0 { return ans }

    ans += "."

    // not all repeating patterns start immediately after the decimal point
    // so we need to store the index position
    posMap := map[int]int{} // remainder -> idx pos
    decimalPart := ""
    for remainder != 0 {
        if idx, ok := posMap[remainder]; ok {
            decimalPart = decimalPart[:idx] + "(" + decimalPart[idx:] + ")"
            return ans + decimalPart
        }

        posMap[remainder] = len(decimalPart)

        remainder *= 10
        digit := remainder / denominator
        decimalPart += strconv.Itoa(digit)
        remainder %= denominator
    }

    return ans + decimalPart
}

// hash map, math, string
// time: O(n)
// space: O(n)