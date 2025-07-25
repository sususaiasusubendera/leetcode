func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	e := make([][]int, n)
	for _, v := range edges {
		e[v[0]] = append(e[v[0]], v[1])
		e[v[1]] = append(e[v[1]], v[0])
	}

	sum := 0
	for _, x := range nums {
		sum ^= x
	}

	res := math.MaxInt32
	var dfs2 func(int, int, int, int) int
	dfs2 = func(x, f, oth, anc int) int {
		son := nums[x]
		for _, y := range e[x] {
			if y == f {
				continue
			}
			son ^= dfs2(y, x, oth, anc)
		}
		if f == anc {
			return son
		}

		res = min(res, calc(oth, son, sum^oth^son))
		return son
	}

	var dfs func(int, int) int
	dfs = func(x, f int) int {
		son := nums[x]
		for _, y := range e[x] {
			if y == f {
				continue
			}
			son ^= dfs(y, x)
		}

		for _, y := range e[x] {
			if y == f {
				dfs2(y, x, son, x)
			}
		}
		return son
	}

	dfs(0, -1)
	return res
}

// hard, skip :')
// editorial
// NOTICE ME SENPAI!!!

func calc(part1, part2, part3 int) int {
	return max(part1, max(part2, part3)) - min(part1, min(part2, part3))
}