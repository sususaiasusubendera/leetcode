func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    start_ := strconv.FormatInt(start - 1, 10)
	finish_ := strconv.FormatInt(finish, 10)
	return calculate(finish_, s, limit) - calculate(start_, s, limit)
}

func calculate(x string, s string, limit int) int64 {
	if len(x) < len(s) {
		return 0
	}
	if len(x) == len(s) {
		if x >= s {
			return 1
		}
		return 0
	}

	suffix := x[len(x) - len(s):]
	var count int64
	preLen := len(x) - len(s)

	for i := 0; i < preLen; i++ {
		digit := int(x[i] - '0')
		if limit < digit {
			count += int64(math.Pow(float64(limit + 1), float64(preLen - i)))
			return count
		}
		count += int64(digit) * int64(math.Pow(float64(limit + 1), float64(preLen - 1- i)))
	}
	if suffix >= s {
		count++
	}
	return count
}

// NOTICE ME SENPAIII!!!!!!!