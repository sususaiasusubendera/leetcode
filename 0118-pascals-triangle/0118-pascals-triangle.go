func generate(numRows int) [][]int {
    result := [][]int{{1}}
    if numRows == 1 {
        return result
    }
    result = append(result, []int{1, 1})
    
    for i := 1; i < numRows-1; i++ {
        row := []int{1}
        for j := 0; j < len(result[i])-1; j++ {
            row = append(row, result[i][j]+result[i][j+1])
        }
        row = append(row, 1)
        result = append(result, row)
    }
    return result
}