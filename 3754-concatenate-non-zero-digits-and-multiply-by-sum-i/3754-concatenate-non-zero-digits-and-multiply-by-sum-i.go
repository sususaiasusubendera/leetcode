func sumAndMultiply(n int) int64 {
    strNum := strconv.Itoa(n)
    sum := int64(0)
    res := []byte{}
    for i := 0; i < len(strNum); i++ {
        if strNum[i] == '0' {
            continue
        }
        sum += int64(strNum[i] - '0')
        res = append(res, strNum[i])
    }

    intRes, _ := strconv.Atoi(string(res))
    return int64(intRes) * sum
}

// math
// time: O(n)
// space: O(n)