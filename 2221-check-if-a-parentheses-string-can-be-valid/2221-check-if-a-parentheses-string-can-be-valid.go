func canBeValid(s string, locked string) bool {
    if len(s) % 2 == 1 {
        return false
    }

    cub := 0 // cui: count unlock bracket
    balance := 0 // balance = locked open bracket - locked close bracket
    // forward pass
    for i := 0; i < len(s); i++ {
        if locked[i] == '0' {
            cub++
        } else if s[i] == '(' {
            balance++
        } else {
            balance--
        }

        if balance + cub < 0 {
            return false
        }
    }

    cub, balance = 0, 0
    // backward pass: balance = locked close bracket - locked open bracket
    for i := len(s)-1; i >= 0; i-- {
        if locked[i] == '0' {
            cub++
        } else if s[i] == ')' {
            balance++
        } else {
            balance --
        }

        if balance + cub < 0 {
            return false
        }
    }

    return true
}