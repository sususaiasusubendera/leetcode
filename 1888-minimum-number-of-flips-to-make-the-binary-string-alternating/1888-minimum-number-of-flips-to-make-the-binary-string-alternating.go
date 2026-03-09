func minFlips(s string) int {
    n := len(s)
    
    s2 := s + s

    o := make([]byte, 2*n) // alternating string, start with 1
    z := make([]byte, 2*n) // alternating string, start with 0
    for i := 0; i < 2*n; i++ {
        if i % 2 == 0 {
            o[i] = '1'
            z[i] = '0'
        } else {
            z[i] = '1'
            o[i] = '0'
        }
    }

    ans := math.MaxInt32
    on, zn := 0, 0
    left, right := 0, 0
    for right < 2 * n {
        if s2[right] != o[right] {
            on++
        }
        if s2[right] != z[right] {
            zn++
        }

        if right - left + 1 > n {
            if s2[left] != o[left] {
                on--
            }
            if s2[left] != z[left] {
                zn--
            }
            left++

            ans = min(ans, min(on, zn))
        }

        right++
    }

    return ans
}

// sliding window, string
// time: O(n)
// space: O(n)