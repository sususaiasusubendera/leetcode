func processStr(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else if s[i] == '#' {
            result = append(result, result...)
        } else if s[i] == '%' {
            result = reverse(result)
        } else {
            result = append(result, s[i])
        }
	}

    return string(result)
}

func reverse(s []byte) []byte {
	temp := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		temp[i] = s[len(s) - i - 1]
	}
	return temp
}

// string
// time: O(nk)
// space: O(n)