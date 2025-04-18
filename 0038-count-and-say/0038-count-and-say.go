func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }

    result := "1"
    for i := 2; i <= n; i++ {
        result = strRLE(parseRLE(result))
    }

    return result
}

// two pointer
// time: O(n)
// space: O(n)

func parseRLE(str string) [][2]int {
    slc := [][2]int{}
    p1 := 0
    for p1 < len(str) {
        countSameNum := 0
        p2 := p1
        for p2 < len(str) && str[p2] == str[p1] {
            p2++
            countSameNum++
        }

        slc = append(slc, [2]int{countSameNum, int(str[p1]-'0')})
        p1 = p2
    }

    return slc
}

func strRLE(slc [][2]int) string {
    res := ""
    for _, v := range slc {
        res += strconv.Itoa(v[0]) + strconv.Itoa(v[1])
    }

    return res
}