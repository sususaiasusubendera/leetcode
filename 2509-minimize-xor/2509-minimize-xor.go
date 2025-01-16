func minimizeXor(num1 int, num2 int) int {
    setBits2 := bits.OnesCount(uint(num2)) // count set bits of num 2
    x := 0
    // use bits from num1, iterate from msb to minimize xor
    for i := 31; i >= 0 && setBits2 > 0; i-- {
        if (num1 & (1 << i)) != 0 { // check if i-th index bit in num1 is '1' (or not '0')
            x |= (1 << i) // set the bit of x in i-th index (with '1')
            setBits2--
        }
    }

    // if more set bits are needed, add from lsb
    for i := 0; setBits2 > 0; i++ {
        if x & (1 << i) == 0 { // check if i-th index bit in x is '0'
            x |= (1 << i)
            setBits2--
        }
    }

    return x
}

// bitwise operation approach
// time: O(1)
// space: O(1)


