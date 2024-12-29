func mergeAlternately(word1 string, word2 string) string {
    result := ""
    turn := 1
    one, two := 0, 0
    steps1, steps2 := len(word1), len(word2)
    for steps1 > 0 || steps2 > 0  {
        if turn == 1 {
            result += string(word1[one])
            if one < len(word1) - 1 {
                one++
            }
            steps1--
            if steps2 > 0 {
                turn = 2
            }
        } else {
            result += string(word2[two])
            if two < len(word2) - 1 {
                two++
            }
            steps2--
            if steps1 > 0 {
                turn = 1
            }
        }
    }
    return result
}