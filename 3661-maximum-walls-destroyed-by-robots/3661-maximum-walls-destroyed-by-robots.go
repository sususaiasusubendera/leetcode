func maxWalls(robots []int, distance []int, walls []int) int {
    n := len(robots)
	r := []Robot{} // robots
	for i, pos := range robots {
		r = append(r, Robot{
			id:  i,
			pos: pos,
			d:   distance[i],
		})
	}

    sort.Slice(r, func(i, j int) bool { return r[i].pos < r[j].pos })
    sort.Slice(walls, func(i, j int) bool { return walls[i] < walls[j] })

    leftCount := make([]int, n)
    rightCount := make([]int, n)
    for i := 0; i < n; i++ {
        p := r[i].pos
        d := r[i].d

        // left
        left := p - d
        if i > 0 {
            left = max(left, r[i-1].pos + 1)
        }
        right := p
        leftCount[i] = countWalls(walls, left, right)

        // right
        left = p
        right = p + d
        if i < n-1 {
            right = min(right, r[i+1].pos - 1)
        }
        rightCount[i] = countWalls(walls, left, right)
    }

    // dp
    dp := make([][2]int, n)
    // base case
    dp[0][0] = leftCount[0]
    dp[0][1] = rightCount[0]
    for i := 1; i < n; i++ {
        // left (overlap for prev right and curr left dp state)
        // MAX of:
        // prev left + curr left
        // prev right - prev count right (basically negate) + MIN(prev count right + curr count left, count walls between prev robot and curr robot)
        dp[i][0] = max(dp[i-1][0] + leftCount[i], dp[i-1][1] - rightCount[i-1] + min(rightCount[i-1] + leftCount[i], countWalls(walls, r[i-1].pos, r[i].pos)))

        // right
        dp[i][1] = max(dp[i-1][0] + rightCount[i], dp[i-1][1] + rightCount[i])
    }

    return max(dp[n-1][0], dp[n-1][1])
}

type Robot struct {
	id  int
	pos int
	d   int
}

func lowerBoundSearch(a []int, t int) int { // bin search
    l, r := 0, len(a)
    for l < r {
        mid := l + ((r - l) / 2)
        if a[mid] >= t {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

func upperBoundSearch(a []int, t int) int { // bin search
    l, r := 0, len(a)
    for l < r {
        mid := l + ((r - l) / 2)
        if a[mid] > t {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

func countWalls(walls []int, l, r int) int {
    left := lowerBoundSearch(walls, l)
    right := upperBoundSearch(walls, r)
    return right - left
}

// array, bin search, dp, sorting
// time: O(n + nlog(n) + mlog(m) + nlog(m))
// space: O(n + m)

// note
// the overlap case in dp is... tricky and...
// what a crazy problem