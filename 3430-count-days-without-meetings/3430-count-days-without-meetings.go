func countDays(days int, meetings [][]int) int {
    events := []Event{}

    for _, meeting := range meetings {
        start, end := meeting[0], meeting[1]
        events = append(events, Event{start, 1})
        events = append(events, Event{end+1, -1})
    }

    sort.Slice(events, func(i, j int) bool {
        return events[i].day < events[j].day
    })

    count, activeMeetings, lastDay := 0, 0, 1
    for _, event := range events {
        if lastDay < event.day && activeMeetings == 0 {
            count += event.day - lastDay
        }
        activeMeetings += event.val
        lastDay = event.day
    }

    if lastDay <= days && activeMeetings == 0 {
        count += days - lastDay + 1
    }

    return count
}

// line sweep
// time: O(n log(n))
// space: O(n)

type Event struct {
    day int
    val int
}