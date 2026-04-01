func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	robots := []Robot{}
	for i := 0; i < n; i++ {
		robots = append(robots, Robot{
			num:    i + 1,
			pos:    positions[i],
			health: healths[i],
			dir:    directions[i],
		})
	}

    // sort by pos
	sort.Slice(robots, func(i, j int) bool { return robots[i].pos < robots[j].pos })

	stack := []Robot{}
	for _, robot := range robots {
		for len(stack) > 0 && stack[len(stack)-1].dir == 'R' && robot.dir == 'L' && stack[len(stack)-1].health < robot.health {
			robot.health--
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1].dir == 'R' && robot.dir == 'L' && stack[len(stack)-1].health > robot.health {
			stack[len(stack)-1].health--
		} else if len(stack) > 0 && stack[len(stack)-1].dir == 'R' && robot.dir == 'L' && stack[len(stack)-1].health == robot.health {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, robot)
		}
	}

    // sort by robot num
    sort.Slice(stack, func(i, j int) bool { return stack[i].num < stack[j].num })

    ans := []int{}
    for _, s := range stack {
        ans = append(ans, s.health)
    }

    return ans
}

type Robot struct {
	num    int
	pos    int
	health int
	dir    byte
}

// array, sorting, stack
// time: O(nlog(n))
// space: O(n)