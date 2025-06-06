func robotWithString(s string) string {
    t := []byte{}
    p := []byte{} // result

    alphabetFreq := make([]int, 26)
    for i := 0; i < len(s); i++ {
        alphabetFreq[s[i]-'a']++
    }

    smallest := 0
    findSmallest(&smallest, alphabetFreq)
    for len(s) > 0 {
        c := s[0] // get s' 1st character
        s = s[1:] // pop s' 1st character
        t = append(t, c) // append c to t

        alphabetFreq[c-'a']-- // decrement the processed character in freq
        // if smallest's freq == 0, find a new smallest
        if alphabetFreq[smallest] == 0 {
            findSmallest(&smallest, alphabetFreq)
        }

        // write t[-1] into p if the condition is met
        for len(t) > 0 && t[len(t)-1] <= byte(smallest+'a') {
            p = append(p, t[len(t)-1])
            t = t[:len(t)-1]
        }

    }

    for len(t) > 0 {
        p = append(p, t[len(t)-1])
        t = t[:len(t)-1]
    }

    return string(p)
}

// greedy, stack
// time: O(n)
// space: O(n)

func findSmallest(smallest *int, alphabetFreq []int) {
    for i := *smallest; i < len(alphabetFreq); i++ {
        if alphabetFreq[i] > 0 {
            *smallest = i
            break
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}