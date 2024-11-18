func groupAnagrams(strs []string) [][]string { // no sort
    m := map[[26]int][]string{}
    for _, str := range strs {
        comp := [26]int{}
        for i := 0; i < len(str); i++ {
            comp[str[i]-'a']++
        }
        m[comp] = append(m[comp], str)
    }

    result := [][]string{}
    for _, v := range m {
        result = append(result, v)
    }
    
    return result
}