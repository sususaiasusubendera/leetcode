func compareVersion(version1 string, version2 string) int {
    ver1, ver2 := strings.Split(version1, "."), strings.Split(version2, ".")

    var l int
    var ver1longer bool
    if len(ver1) < len(ver2) {
        l = len(ver1)
        ver1longer = false
    } else {
        l = len(ver2)
        ver1longer = true
    }

    idx := 0
    for idx < l {
        if strToInt(ver1[idx]) < strToInt(ver2[idx]) {
            return -1
        } else if strToInt(ver1[idx]) > strToInt(ver2[idx]) {
            return 1
        }

        idx++
    }

    if ver1longer {
        for idx < len(ver1) {
            if strToInt(ver1[idx]) > 0 {
                return 1
            }
            idx++
        }
    } else {
        for idx < len(ver2) {
            if strToInt(ver2[idx]) > 0 {
                return -1
            }
            idx++
        }
    }

    return 0
}

// string, two pointer
// time: O(n + m)
// space: O(n + m)

func strToInt(s string) int {
    num, _ := strconv.Atoi(s)
    return num
}