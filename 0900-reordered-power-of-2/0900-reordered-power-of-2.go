func reorderedPowerOf2(n int) bool {
    powSet := map[string]bool{}
    for num := 1; num <= 1e9; num <<= 1 {
        strNum := strconv.Itoa(num)
        strNum = sortStr(strNum)
        powSet[strNum] = true
    }

    sortedN := sortStr(strconv.Itoa(n))
    return powSet[sortedN]
}

// time: O(1)
// space: O(1)

func sortStr(s string) string {
    b := []byte(s)
    sort.Slice(b, func(i, j int) bool {
        return b[i] < b[j]
    })

    return string(b)
}