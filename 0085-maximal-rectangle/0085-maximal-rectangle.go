func maximalRectangle(matrix [][]byte) int {
	ans := make([]int, len(matrix[0]))
	mx := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				ans[j] = 0
			} else {
				ans[j]++
			}

		}
		mx = Max(mx, largestRectangleArea(ans, len(ans)))
	}
	return mx
}

func largestRectangleArea(heights []int, n int) int {
	nsl := make([]int, len(heights))
	stk := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			nsl[i] = -1
		} else {
			nsl[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	nsr := make([]int, len(heights))
	stk = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			nsr[i] = len(heights)
		} else {
			nsr[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		width := nsr[i] - nsl[i] - 1
		area := heights[i] * width
		maxArea = int(math.Max(float64(maxArea), float64(area)))
	}

	return maxArea
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// NOTICE ME SENPAI
// array, dp, matrix, mono stack, stack
// solution from solutions @sandy0981