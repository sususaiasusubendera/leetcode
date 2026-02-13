func longestBalanced(s string) int {
    res := []int{
        solve1(s, 'a'),
        solve1(s, 'b'),
        solve1(s, 'c'),
        solve2(s, 'a', 'b'),
        solve2(s, 'a', 'c'),
        solve2(s, 'b', 'c'),
        solve3(s),
    }
    
    ans := 0
    for _, r := range res {
        ans = max(ans, r)
    }
    return ans
}

// hash map, prefix sum, string
// ref: tdkkdt
// time: O(n)
// space: O(n)

func solve1(s string, t byte) int { // t: target
    res, count := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == t {
            count++
            res = max(res, count)
            continue
        }
        count = 0
    }
    return res
}

func solve2(s string, t1, t2 byte) int {
    res, c1, c2 := 0, 0, 0
    prev := map[int]int{}
    for i := 0; i < len(s); i++ {
        if s[i] != t1 && s[i] != t2 {
            prev = map[int]int{}
            c1, c2 = 0, 0
            continue
        }

        if s[i] == t1 {
            c1++
        } else {
            c2++
        }

        if c1 == c2 {
            l := c1 + c2 // length
            res = max(res, l)
        } else {
            diff := c1 - c2
            if p, ok := prev[diff]; ok {
                l := i - p
                res = max(res, l)
            } else {
                prev[diff] = i
            }
        }
    }
    return res
}

type Pair struct {
    x, y int
}

func solve3(s string) int {
    res, ca, cb, cc := 0, 0, 0, 0
    prev := map[Pair]int{}
    for i := 0; i < len(s); i++ {
        if s[i] == 'a' {
            ca++
        } else if s[i] == 'b' {
            cb++
        } else {
            cc++
        }

        if ca == cb && cb == cc {
            res = i + 1
        } else {
            diff := Pair{ca-cb, cb-cc}
            if p, ok := prev[diff]; ok {
                l := i - p
                res = max(res, l)
            } else {
                prev[diff] = i
            }
        }
    }
    return res
}