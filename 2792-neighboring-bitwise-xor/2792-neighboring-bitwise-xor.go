func doesValidArrayExist(derived []int) bool {
    original := make([]int, len(derived))

    original[0] = 0 // assumption original[0] = 0
    for i := 1; i < len(original); i++ {
        original[i] = original[i-1] ^ derived[i-1]
    }
    if original[len(original)-1] ^ derived[len(original)-1] == 0 {
        return true
    }

    original[0] = 1 // assumption original[0] = 1
    for i := 1; i < len(original); i++ {
        original[i] = original[i-1] ^ derived[i-1]
    }
    if original[len(original)-1] ^ derived[len(original)-1] == 1 {
        return true
    }

    // no valid loop check
    return false
}

// time: O(n)
// space: O(n)

// xor (^) properties related
// 0 ^ 0 = 0
// 1 ^ 0 = 1
// 0 ^ 1 = 1
// 1 ^ 1 = 0
// a ^ b = c means b = a ^ c or a = b ^ c