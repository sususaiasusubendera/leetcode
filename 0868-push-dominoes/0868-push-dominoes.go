func pushDominoes(dominoes string) string {
    ans := []byte(dominoes)
    
    dir := byte('L')
    l := 0
    for r:=range ans{
        if ans[r] == '.'{
            continue
        }
        
        if ans[r] == dir{
            for ;l<=r;l++{
                ans[l]=dir
            } 
        }else{
            if ans[r] == 'L' && dir == 'R'{
                rr := r-1
                ll := l
                for ll < rr {
                    ans[ll]='R'
                    ans[rr]='L'
                    ll++
                    rr--
                }
            }
            
            l=r+1
        }
        
        dir=ans[r]
    }
    
    if dir == 'R'{
        for ;l<len(ans);l++{
            ans[l]='R'
        }
    }
    
    return string(ans)
}

// NOTICE ME SENPAI