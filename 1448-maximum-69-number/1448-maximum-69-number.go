func maximum69Number (num int) int {
    strNum := strconv.Itoa(num)
    strResNum := []byte{}
    change := false
    for i := 0; i < len(strNum); i++ {
        if strNum[i] == '6' && !change {
            strResNum = append(strResNum, '9')
            change = true
            continue
        }

        strResNum = append(strResNum, strNum[i])
    }

    resNum, _ := strconv.Atoi(string(strResNum))
    return resNum
}

// greedy
// time: O(n)
// space: O(n)