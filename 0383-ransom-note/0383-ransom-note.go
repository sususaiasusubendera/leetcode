func canConstruct(ransomNote string, magazine string) bool {
    m := map[byte]int{}
    for i := 0; i < len(magazine); i++ {
        m[magazine[i]] += 1
    }

    for i := 0; i < len(ransomNote); i++ {
        if _, exist := m[ransomNote[i]]; !exist || m[ransomNote[i]] == 0 {
            return false
        } else {
            m[ransomNote[i]] -= 1
        }
    }

    return true
}