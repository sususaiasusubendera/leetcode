func isOneBitCharacter(bits []int) bool {
    idx := 0
    for idx <= len(bits)-2 {
        if bits[idx] == 1 {
            idx += 2
        } else if bits[idx] == 0 {
            idx += 1
        }
    }

    if idx < len(bits) { return true }
    return false
}

// array
// time: O(n)
// space: O(1)