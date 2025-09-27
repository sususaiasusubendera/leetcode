func calcArea(a, b, c []int) int {
	return a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1])
}

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	result := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				area := calcArea(points[i], points[j], points[k])
				result = max(result, area, -area)
			}
		}
	}
	return 0.5 * float64(result)
}

// solution