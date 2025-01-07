func shiftingLetters(s string, shifts [][]int) string {
    // initiate diff slice with size len(s)+1 (for difference array approach)
    diff := make([]int, len(s)+1)

    // the difference array approach
    for _, shift := range shifts {
        start, end, direction := shift[0], shift[1], shift[2]
        // mark the start index and end+1 index of the diff according to the direction of shift
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
        shifted := (int(s[i]-'a') + currentShift) % 26 // shift the s[i] according to the currentShift
        if shifted < 0 { // case for backward shift past 'a'
            shifted += 26
        }
        result[i] = byte(shifted) + 'a'
    }

    return string(result)
}

// differecne array approach