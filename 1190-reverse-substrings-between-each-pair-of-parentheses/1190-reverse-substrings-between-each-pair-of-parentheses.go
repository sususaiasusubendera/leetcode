func reverseParentheses(s string) string {
    openParanthesesIdx := []int{} // stack
    result := []rune{}
    for _, c := range s {
        if c == '(' {
            openParanthesesIdx = append(openParanthesesIdx, len(result))
        } else if c == ')' {
            idx := openParanthesesIdx[len(openParanthesesIdx)-1]
            openParanthesesIdx = openParanthesesIdx[:len(openParanthesesIdx)-1] // pop
            reverse(result, idx)
        } else { // character
            result = append(result, c)
        }
    }

    return string(result)
}

func reverse(s []rune, startIdx int) {
    left, right := startIdx, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}

// stack, string
// time: O(n^2)
// space: O(n)