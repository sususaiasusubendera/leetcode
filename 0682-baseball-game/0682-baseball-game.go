func calPoints(operations []string) int {
    stack := []int{}
    total := 0
    for i := 0; i < len(operations); i++ {
        switch operations[i] {
            case "+":
                score := stack[len(stack)-1] + stack[len(stack)-2]
                stack = append(stack, score)
                total += score
            case "D":
                score := stack[len(stack)-1] * 2
                stack = append(stack, score)
                total += score
            case "C":
                score := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                total -= score
            default:
                n, _ := strconv.Atoi(operations[i])
                stack = append(stack, n)
                total += n
                
        }
    }
    return total
}