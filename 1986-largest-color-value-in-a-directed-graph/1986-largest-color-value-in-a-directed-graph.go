func largestPathValue(colors string, edges [][]int) int {
    var n int = len(colors)
    adjList := make([][]int, n)
    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
    }
    var ans int = math.MinInt
    visited := make([]bool, n) 
    currPath := make([]bool, n)
    count := make([][]int, n)
    for i := 0; i < n; i++ {
        count[i] = make([]int, 26)
    }
    for i := 0; i < n; i++ {
        dfs(i, colors, adjList, visited, currPath, count, &ans)
    }
    return ans
}

// NOTICE ME SENPAI!!!

func dfs(curr int, colors string, adjList [][]int, visited, currPath []bool, count [][]int, ans *int) {
    if currPath[curr] || *ans == -1 {
        *ans = -1
        return
    }
    if visited[curr] {return}
    visited[curr] = true
    currPath[curr] = true
    for _, next := range adjList[curr] {
        dfs(next, colors, adjList, visited, currPath, count, ans)
        if *ans == -1 {return}
        for i := 0; i < 26; i++ {
            if count[next][i] > count[curr][i] {
                count[curr][i] = count[next][i]
            }
        }
    }
    count[curr][colors[curr] - 'a']++
    if count[curr][colors[curr] - 'a'] > *ans {*ans = count[curr][colors[curr] - 'a']}
    currPath[curr] = false
}