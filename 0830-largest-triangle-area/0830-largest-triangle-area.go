func largestTriangleArea(points [][]int) float64 {
	max := 0.0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				area := triangleArea(points[i], points[j], points[k])
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

// math
// time: O(n^3)
// space: O(1)

func triangleArea(a, b, c []int) float64 {
	xa, ya := a[0], a[1]
	xb, yb := b[0], b[1]
	xc, yc := c[0], c[1]

	p1 := float64(xa) * (float64(yb) - float64(yc))
	p2 := float64(xb) * (float64(yc) - float64(ya))
	p3 := float64(xc) * (float64(ya) - float64(yb))

	return 0.5 * abs(p1+p2+p3)
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
} 