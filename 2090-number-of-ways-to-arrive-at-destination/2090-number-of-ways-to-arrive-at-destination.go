func countPaths(n int, roads [][]int) int {
    mod:=int(1e9)+7
    dp:=make([][][2]int,n)
    for i:=range n{
        dp[i]=make([][2]int,n)
        for j:=range n{
            if i==j{
                dp[i][j][1]=1
            }else{
                dp[i][j][0]=1e12
            }
        }
    }

    for _,r:=range roads{
        u,v,t:=r[0],r[1],r[2]
        dp[u][v][0]=t
        dp[u][v][1]=1
        dp[v][u][0]=t
        dp[v][u][1]=1
    }

    for mid:=range n{
        for s:=range n{
            for d:=range n{
                if s==mid || d==mid{
                    continue
                }

                newT:=dp[s][mid][0]+dp[mid][d][0]
                if newT<dp[s][d][0]{
                    dp[s][d][0]=newT
                    dp[s][d][1]=(dp[s][mid][1]*dp[mid][d][1])%mod
                }else if newT==dp[s][d][0]{
                    dp[s][d][1]=(dp[s][d][1]+dp[s][mid][1]*dp[mid][d][1])%mod
                }
            }
        }
    }
    return dp[0][n-1][1]
}

// NOTICE ME SENPAI!!!