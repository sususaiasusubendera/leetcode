func shortestCommonSupersequence(str1 string, str2 string) string {
  lcsStr := LCS(str1, str2)
  res := ""
  i, j := 0, 0
  for k := range lcsStr {
      for str1[i] != lcsStr[k] {
          res += string(str1[i])
          i++
      }

      for str2[j] != lcsStr[k] {
          res += string(str2[j])
          j++
      }
      res += string(lcsStr[k])
      i++
      j++
  }
  return res + string(str1[i:]) + string(str2[j:])

}

func LCS(a, b string) string {
    dp := make([][]string, len(a)+1)
    for i := range dp {
        dp[i] = make([]string, len(b)+1)
    }

    for i:=1; i<=len(a); i++ {
        for j:=1; j<=len(b); j++ {
            if a[i-1] == b[j-1] {
                dp[i][j] = dp[i-1][j-1] + string(a[i-1])
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[len(a)][len(b)]
}


func max(a, b string) string {
    if len(a) > len(b) {
        return a
    }
    return b
}