type Bank struct {
    accounts []int64
}


func Constructor(balance []int64) Bank {
    bank := Bank{accounts: make([]int64, len(balance)+1)}
    for i := 0; i < len(balance); i++ {
        bank.accounts[i+1] = balance[i]
    }

    return bank
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1 < 1 || account1 > len(this.accounts)-1 || account2 < 1 || account2 > len(this.accounts)-1 {
        return false
    }

    if money > this.accounts[account1] {
        return false
    }

    this.accounts[account1] -= money
    this.accounts[account2] += money

    return true
}


func (this *Bank) Deposit(account int, money int64) bool {
    if account < 1 || account > len(this.accounts)-1 {
        return false
    }

    this.accounts[account] += money

    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if account < 1 || account > len(this.accounts)-1 {
        return false
    }

    if money > this.accounts[account] {
        return false
    }

    this.accounts[account] -= money

    return true
}

// array, design
// time:
// - constructor: O(n)
// - transfer: O(1)
// - deposit: O(1)
// - withdraw: O(1)
// space: O(n)

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */