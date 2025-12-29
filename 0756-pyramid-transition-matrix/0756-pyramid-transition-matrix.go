func pyramidTransition(bottom string, allowed []string) bool {
    patterns := map[string][]byte{}
    for _, rule := range allowed {
        patterns[rule[:2]] = append(patterns[rule[:2]], rule[2])
    }

    return canBuild(bottom, "", 0, patterns)
}

func canBuild(curr string, next string, idx int, patterns map[string][]byte) bool {
    if len(curr) == 1 {
        return true
    }

    if idx == len(curr) - 1 {
        return canBuild(next, "", 0, patterns)
    }

    base := curr[idx:idx+2]
    if tops, ok := patterns[base]; ok {
        for _, top := range tops {
            if canBuild(curr, next+string(top), idx+1, patterns) {
                return true
            }
        }
    }

    return false
}

// backtracking, hash map, string
// time: O(k^n)
// space: O(n) 