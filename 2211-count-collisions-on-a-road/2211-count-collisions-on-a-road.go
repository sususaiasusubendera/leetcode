func countCollisions(directions string) int {
    count := 0
    stack := []rune{}
    for _, d := range directions {
        if len(stack) > 0 {
            if stack[len(stack)-1] == 'R' {
                if d == 'L' {
                    count += 2
                    stack = stack[:len(stack)-1] // pop
                    for len(stack) > 0 && stack[len(stack)-1] == 'R' {
                        count += 1 // R collides with S (RL -> S)
                        stack = stack[:len(stack)-1] // pop
                    }
                    stack = append(stack, 'S')
                } else if d == 'S' {
                    for len(stack) > 0 && stack[len(stack)-1] == 'R' {
                        count += 1 // R collides with S
                        stack = stack[:len(stack)-1] // pop
                    }
                    stack = append(stack, 'S')
                } else { // R
                    stack = append(stack, 'R')
                }
            } else if stack[len(stack)-1] == 'S' {
                if d == 'L' {
                    count += 1
                    stack = append(stack, 'S')
                } else if d == 'R' {
                    stack = append(stack, 'R')
                } else { //S
                    stack = append(stack, 'S')
                }
            } else if stack[len(stack)-1] == 'L' {
                stack = append(stack, d)
            }
            continue
        }

        stack = append(stack, d)
    }

    return count
}

// simulation, stack, string
// time: O(n)
// space: O(n)
