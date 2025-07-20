func removeSubfolders(folders []string) []string {
    m := make(map[string]interface{}, len(folders))
    for _, folder := range folders {
        m[folder] = nil
    }

    result := make([]string, 0, len(folders))
    for _, folder := range folders {
        isSub := false
        for i := range folder[:len(folder)-1] {
            if folder[i] != '/' {
                continue
            }

            if _, exists := m[folder[:i]]; exists {
                isSub = true
                break
            }
        }

        if isSub {
            continue
        }

        result = append(result, folder)
    }

    return result
}

// string, map
// time: O(nL)
// space: O(n)