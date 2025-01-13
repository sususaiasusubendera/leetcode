func fizzBuzz(n int) []string {
    answer := make([]string, n)
    for i := 1; i <= n; i++ {
        switch {
            case i % 3 == 0 && i % 5 == 0:
                answer[i-1] = "FizzBuzz"
            case i % 3 == 0:
                answer[i-1] = "Fizz"
            case i % 5 == 0:
                answer[i-1] = "Buzz"
            default:
                answer[i-1] = strconv.Itoa(i)
        }
    }
    return answer
}