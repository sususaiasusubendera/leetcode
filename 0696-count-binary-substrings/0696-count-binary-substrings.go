func countBinarySubstrings(s string) int {
    ans := 0
    for i := 0; i < len(s); i++ {
        count := 0
        curr := s[i]
        j := i
        for j < len(s) && s[j] == curr {
            count++
            j++
        }
        
        found := false
        for j < len(s) && s[j] != curr {
            count--
            j++
            if count == 0 {
                found = true
                break
            }
        }

        if found { ans++ }
    }

    return ans
}

// string, two pointers
// time: O(n)
// space: O(1)