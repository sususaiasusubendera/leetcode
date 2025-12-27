func mostBooked(n int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })

    room := make([]int, n)
    freq := make([]int, n)
    for _, meeting := range meetings {
        start, end := meeting[0], meeting[1]
        earliestRoom := 0
        earliestTime := room[0]
        take := false
        for r := 0; r < n; r++ {
            // determine the earliest room
            if room[r] < earliestTime {
                earliestTime = room[r]
                earliestRoom = r
            }

            if room[r] <= start {
                take = true
                room[r] = end
                freq[r]++
                break
            }
        }

        // take the earliest room if all rooms are full
        if !take {
            freq[earliestRoom]++
            room[earliestRoom] += (end - start)
        }
    }

    best := 0
    for i := 1; i < n; i++ {
        if freq[i] > freq[best] {
            best = i
        }
    }

    return best
}

// array, simulation, sorting
// time: O(mlog(m) + mn)
// space: O(n)