func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	
	if numRows == 1 {
		return result
	}
	
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		row := []int{1}
		for j := 1; j < i; j++ {
	   		row = append(row, result[i-1][j-1]+result[i-1][j])
		}
		row = append(row, 1)
	
		result[i] = row  
	}

    return result
}

// time: O(n^2)
// space: O(n^2)
