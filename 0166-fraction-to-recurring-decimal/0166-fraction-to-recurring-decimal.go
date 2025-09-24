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

    // remainder part
    remainder := numerator % denominator
    if remainder == 0 { return ans }

    ans += "."

    posMap := map[int]int{}
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

// temp