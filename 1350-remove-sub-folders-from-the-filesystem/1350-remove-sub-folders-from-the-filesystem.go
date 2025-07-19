func removeSubfolders(folders []string) []string {
    folderSet := make(map[string]any, len(folders))
    for _, folder := range folders {
        folderSet[folder] = nil
    }

    res := make([]string, 0, len(folders))
    for _, folder := range folders {
        isSub := false
        for i := range folder[:len(folder)-1] {
            if folder[i] != '/' {
                continue
            }
            if _, ok := folderSet[folder[:i]]; ok {
                isSub = true
                break
            }
        }
        if isSub {
            continue
        }
        res = append(res, folder)
    }
    return res
}

// notice me senpai!