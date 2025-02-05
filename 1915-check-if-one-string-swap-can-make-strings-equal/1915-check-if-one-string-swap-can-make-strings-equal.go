func areAlmostEqual(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }

    // find all index that have different string in s1 and s2
    diffIndex := []int{}
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diffIndex = append(diffIndex, i)
        }
    }

    // if len diffIndex != 2, can't perform exactly one string swap
    if len(diffIndex) != 2 {
        return false
    }

    // check each set of caracters (diff index (2 index))
    return (s1[diffIndex[0]] == s2[diffIndex[1]]) && (s1[diffIndex[1]] == s2[diffIndex[0]])
}

// time: O(n)
// space: O(1)