func wordSubsets(words1 []string, words2 []string) []string {
    mcf2 := make([]int, 26) // max character frequency words2
    cc := make(map[int]struct{}) // characters that need to be check
    // count max character frequency for words2
    for i := 0; i < len(words2); i++ {
        cf2 := countCharFreq(words2[i]) // character frequency
        for j := 0; j < 26; j++ {
            if cf2[j] > mcf2[j] {
                mcf2[j] = cf2[j]
                if _, exist := cc[j]; !exist {
                    cc[j] = struct{}{}
                }
            }
        }
    }

    // check if any universal word is in words1 through character frequency checking with mcf
    result := []string{}
    for i := 0; i < len(words1); i++ {
        cf1 := countCharFreq(words1[i]) // character frequency word of words1
        if isSubset(cf1, mcf2, cc) {
            result = append(result, words1[i])
        }
    }
    return result
}
// character frequency naive solution

// UTILS
func countCharFreq(chars string) []int {
    cf := make([]int, 26)
    for _, c := range chars {
        idx := c - 'a'
        cf[idx]++
    }
    return cf
}

func isSubset(cfSet []int, cfSubset []int, cCheck map[int]struct{}) bool {
    for idx := range cCheck {
        if cfSubset[idx] > cfSet[idx] {
            return false
        }
    }
    return true
}