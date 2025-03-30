func partitionLabels(s string) []int {
    lastOccurrence := make([]int, 26)
    for i, ch := range s {
        lastOccurrence[ch-'a'] = i
    }
    
    var result []int
    start, end := 0, 0
    
    for i, ch := range s {
        if lastOccurrence[ch-'a'] > end {
            end = lastOccurrence[ch-'a']
        }
        
        if i == end {
            result = append(result, end-start+1)
            start = end + 1
        }
    }
    
    return result
}

// notice me senpai!