func closestMeetingNode(edges []int, node1 int, node2 int) int {
    n := len(edges)

    dihstance := func(start int)[]int{
        dist := make([]int, n)
        for i := range dist{
            dist[i]= -1
        }
        d := 0
        for cur := start; cur != -1 && dist[cur]==-1; cur=edges[cur]{
            dist[cur] = d
            d++
        }
        return dist
    }
    dist1:= dihstance(node1)
    dist2:= dihstance(node2)
    mDist := int(1e9)
    res := -1
    for i:=0; i < n; i++{
        if dist1[i]!=-1 && dist2[i]!=-1{
            maxDist:=max(dist1[i], dist2[i])
            if maxDist < mDist{
                mDist = maxDist
                res=i
            }
        }
    }
    return res
}

// NOTICE ME SENPAI!!!