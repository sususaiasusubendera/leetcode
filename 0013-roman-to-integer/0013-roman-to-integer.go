func romanToInt(s string) int {
    roman := map[string]int {
        "I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
    }

    val := 0
    for i, v := range s {
        val += roman[string(v)]
        if i != 0 {
            if roman[string(s[i-1])] < roman[string(v)] {
                val -= 2 * roman[string(s[i-1])]
            }
        }
    }
    return val
}