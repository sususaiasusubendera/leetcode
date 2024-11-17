func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    rows := make([][]byte, numRows)
    var up bool
    curRow := 0
    for i := 0; i < len(s); i++ {
        rows[curRow] = append(rows[curRow], s[i])
        if curRow == 0 {
            up = false
        } else if curRow == numRows-1 {
            up = true
        }
        if !up {
            curRow++
        } else {
            curRow--
        }
    }

    var result string
    for i := 0; i < numRows; i++ {
        result += string(rows[i])
    }

    return result
}