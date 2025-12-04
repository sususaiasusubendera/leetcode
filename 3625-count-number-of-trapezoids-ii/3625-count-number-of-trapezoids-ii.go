func countTrapezoids(points [][]int) int {
	n := len(points)

	lineFreq := map[Line]int{} // map line -> segment freq on that line
	segments := []Segment{}
	for i := 0; i < n-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x2 - x1
			dy := y2 - y1

			// normalize slope using GCD (reducing fraction)
			g := gcd(dy, dx)
			dy /= g
			dx /= g

			// make dx >= 0 for consistency (unique slope)
			if dx < 0 {
				dx, dy = -dx, -dy
			}

			// handle vertical line (dx == 0)
			if dx == 0 {
				dy = 1
				c := x1 // constant = x

				l := Line{Slope{dy, dx}, c}
				lineFreq[l]++
				segments = append(segments, Segment{i, j, l})
				continue
			}

			// compute the line constant (c)
			// it uniquely identifies the lines with the same slope
			c := dy*x1 - dx*y1

			l := Line{Slope{dy, dx}, c}
			lineFreq[l]++
			segments = append(segments, Segment{i, j, l})
		}
	}

	// group segment freq by slope (parallel lines)
	slopeGroup := map[Slope][]int{}
	for l, f := range lineFreq {
		slopeGroup[l.m] = append(slopeGroup[l.m], f)
	}

	// trapezoids: pick segments from different parallel lines
	trapezoids := 0
	for _, arr := range slopeGroup {
		sum := 0
		for _, f := range arr {
			trapezoids += sum * f
			sum += f
		}
	}

	// parallelograms: segments with the same midpoint
	midLineFreq := map[Mid]map[Line]int{} // map mid -> (map line -> segment freq)
	midTotal := map[Mid]int{} // map midpoint -> segment freq with the same midpoint 
	for _, s := range segments {
		i, j := s.i, s.j
		sx := points[i][0] + points[j][0]
		sy := points[i][1] + points[j][1]
		mid := Mid{sx, sy}

		if midLineFreq[mid] == nil {
			midLineFreq[mid] = map[Line]int{}
		}

		midLineFreq[mid][s.l]++
		midTotal[mid]++
	}

    // remove parallelograms from trapezoids count (to avoid double-counting)
    // parallelogram has 2 parallel pairs -> double-counting
	parallelograms := 0
	for mid, total := range midTotal {
		if total < 2 {
			continue
		}

		// all segment pairs that have same midpoint
		totalPairs := total * (total - 1) / 2

		// remove collinear pairs (same line)
		sameLinePairs := 0
		for _, cnt := range midLineFreq[mid] {
			if cnt >= 2 {
				sameLinePairs += cnt * (cnt - 1) / 2
			}
		}

		parallelograms += (totalPairs - sameLinePairs)
	}

	return trapezoids - parallelograms
}

type Slope struct {
	dy int
	dx int
}

type Line struct {
	m Slope
	c int // line constant (c in y = mx + c): dy*x - dx*y
}

type Segment struct {
	i, j int
	l    Line
}

type Mid struct {
	sx int // sum x
	sy int // sum y
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// array, geometry, hash map, math
// time: O(n^2)
// space: O(n^2)