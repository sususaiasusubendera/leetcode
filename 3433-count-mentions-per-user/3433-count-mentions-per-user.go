func countMentions(numberOfUsers int, events [][]string) []int {
    sort.Slice(events, func(i, j int) bool {
        ti, _ := strconv.Atoi(events[i][1])
        tj, _ := strconv.Atoi(events[j][1])

        if ti == tj {
            var messageI, messageJ int // MESSAGE: 2, OFFLINE: 1
            if events[i][0] == "MESSAGE" {
                messageI = 2
            } else { // OFFLINE
                messageI = 1
            }

            if events[j][0] == "MESSAGE" {
                messageJ = 2
            }  else { // OFFLINE
                messageJ = 1
            }

            return messageI < messageJ
        }

        return ti < tj
    })

	mentions := make([]int, numberOfUsers)
	offlineUntil := make([]int, numberOfUsers) // stores user's offline timestamp
	for _, e := range events {
		ts, _ := strconv.Atoi(e[1]) // timestamp
		if e[0] == "MESSAGE" {
			s := strings.Fields(e[2])
			if len(s) == 1 {
				if s[0] == "ALL" {
					for i := range mentions {
						mentions[i]++
					}
				} else if s[0] == "HERE" {
					for id, ou := range offlineUntil {
                        if ts >= ou {
                            mentions[id]++
                        } 
                    }
				} else { // s[0] == id
					id := parseId(s[0])
					mentions[id]++
				}
			} else {
				for _, v := range s {
					id := parseId(v)
					mentions[id]++
				}
			}
		} else if e[0] == "OFFLINE" {
			id, _ := strconv.Atoi(e[2])
			offlineUntil[id] = ts + 60
		}
	}

	return mentions
}

func parseId(s string) int {
	var d strings.Builder // digit
	for i := 2; i < len(s); i++ {
		d.WriteByte(s[i])
	}

	id, _ := strconv.Atoi(d.String())
	return id
}

// array, simulation, sorting
// time: O(Mlog(M) + MN)
// space: O(M + N)