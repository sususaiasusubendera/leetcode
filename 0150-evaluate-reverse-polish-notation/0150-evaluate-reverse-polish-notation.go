func evalRPN(tokens []string) int {
    stack := []int{}
    for i := 0; i < len(tokens); i++ {
        switch tokens[i] {
            case "+", "-", "*", "/":
                var result int
                if tokens[i] == "+" {
                    result = stack[len(stack)-2] + stack[len(stack)-1]
                } else if tokens[i] == "-" {
                    result = stack[len(stack)-2] - stack[len(stack)-1]
                } else if tokens[i] == "*" {
                    result = stack[len(stack)-2] * stack[len(stack)-1]
                } else {
                    result = stack[len(stack)-2] / stack[len(stack)-1]
                }
                stack = stack[:len(stack)-2]
                stack = append(stack, result)
            default:
                num, _ := strconv.Atoi(tokens[i])
                stack = append(stack, num)
        }
    }
    return stack[0] 
}