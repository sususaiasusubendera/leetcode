func shiftingLetters(s string, shifts [][]int) string {
    diff := make([]int, len(s)+1)

    for _, shift := range shifts {
        start, end, direction := shift[0], shift[1], shift[2]
        if direction == 1 {
            diff[start]++
            diff[end+1]--
        } else {
            diff[start]--
            diff[end+1]++
        }
    }

    currentShift := 0
    result := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        currentShift += diff[i]
        shifted := (int(s[i]-'a') + currentShift) % 26
        if shifted < 0 {
            shifted += 26
        }
        result[i] = byte(shifted) + 'a'
    }

    return string(result)
}

// temp