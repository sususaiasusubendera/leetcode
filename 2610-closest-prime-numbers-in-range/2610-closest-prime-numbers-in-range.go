func closestPrimes(left int, right int) []int {
    isPrime := make([]bool, right+1)
    for i := 2; i < len(isPrime); i++ {
        isPrime[i] = true
    }

    for i := 2; i*i <= right; i++ {
        if isPrime[i] {
            for j := i*i; j <= right; j += i {
                isPrime[j] = false
            }
        }
    }

    primes := []int{}
    for i := left; i <= right; i++ {
        if isPrime[i] {
            primes = append(primes, i)
        }
    }

    minDiff := math.MaxInt
    num1, num2 := -1, -1
    for i := 0; i < len(primes)-1; i++ {
        diff := primes[i+1] - primes[i]
        if diff < minDiff {
            minDiff = diff
            num1, num2 = primes[i], primes[i+1]
        }
    }
    
    return []int{num1, num2}
}

// sieve of eratosthenes approach
// time: O(n log log n)
// space: O(n)