func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    userLang := make([]map[int]bool, len(languages)+1) // slice that represent users with langs that they know (in map) 
    for i := 0; i < len(userLang)-1; i++ {
        userLang[i+1] = map[int]bool{} // init map (int -> bool) for each user
        for _, l := range languages[i] { // add user's known langs to the map
            userLang[i+1][l] = true
        }
    }

    userLearn := map[int]bool{} // set of users that need to learn new langs
    for _, f := range friendships {
        u, v := f[0], f[1]
        canCommunicate := false
        for k := range userLang[u] {
            if userLang[v][k] { // if there is 1 lang that u and v know, ok
                canCommunicate = true
                break
            }
        }
        if !canCommunicate { // add u and v to userLearn if they can't communicate each other
            userLearn[u] = true
            userLearn[v] = true
        }
    }

    langCount := map[int]int{} // count language that user (from userLearn) know
    for user := range userLearn {
        for lang := range userLang[user] {
            langCount[lang]++
        }
    }

    max := 0
    for _, c := range langCount {
        if c > max { max = c }
    }

    return len(userLearn) - max
}

// greedy, hash map
// time: O(nm + fm + nm + l)
// space: O(nm + n + l)