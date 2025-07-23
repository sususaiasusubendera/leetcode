func maximumGain(s string, x int, y int) int {
    if x < y {
        x, y = y, x
        s = swapAB(s)
    }

    stackA := []byte{}
    score := 0
    for i := 0; i < len(s); i++ {
        if len(stackA) > 0 && stackA[len(stackA)-1] == 'a' && s[i] == 'b' {
            stackA = stackA[:len(stackA)-1] // pop 'a'
            score += x
        } else {
            stackA = append(stackA, s[i])
        }
    }

	stackB := []byte{}
	for i := 0; i < len(stackA); i++ {
		if len(stackB) > 0 && stackB[len(stackB)-1] == 'b' && stackA[i] == 'a' {
			stackB = stackB[:len(stackB)-1] // pop 'b'
			score += y
		} else {
			stackB = append(stackB, stackA[i])
		}
	}

	return score
}

// greedy, stack
// time: O(n)
// space: O(n)

func swapAB(s string) string {
    bytes := []byte(s)
    for i := range bytes {
        if bytes[i] == 'a' {
            bytes[i] = 'a' + 100 // temp
        } else if bytes[i] == 'b' {
            bytes[i] = 'a'
        }
    }

    for i := range bytes {
        if bytes[i] == 'a' + 100 {
            bytes[i] = 'b'
        }
    }

    return string(bytes)
}