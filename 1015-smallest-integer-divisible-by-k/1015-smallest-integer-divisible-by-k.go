func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 || k%10 == 0 { return -1 }

    // tip: store the modulo value, not the repunit number, to avoid overflow
    rem := 0 // remainder
    for length := 1; length <= k; length++ {
        rem = ((rem * 10) + 1) % k
        if rem == 0 { return length }
    } 

    return -1
}