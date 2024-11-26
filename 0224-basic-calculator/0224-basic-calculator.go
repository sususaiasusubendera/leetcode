func calculate(s string) int {
    stack := []int{}
    currNum, currRes, sign := 0, 0, 1
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            continue
        }

        switch s[i] {
            case '+', '-':
                currRes = currRes + (currNum * sign)
                if s[i] == '+' {
                    sign = 1
                } else {
                    sign = -1
                }
                currNum = 0
            case '(': // push the currRes before '('
                stack = append(stack, currRes)
                stack = append(stack, sign)
                currRes = 0
                sign = 1
            case ')': // calculate, pop, calculate
                currRes = currRes + currNum * sign
                currNum = 0
                sign = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                currRes = stack[len(stack)-1] + currRes * sign
                stack = stack[:len(stack)-1]
            default:
                currNum = currNum*10 + int(s[i] - '0')
        }
    }

    // check the last element
    if currNum != 0 {
        currRes = currRes + (currNum * sign)
    }

    return currRes
}