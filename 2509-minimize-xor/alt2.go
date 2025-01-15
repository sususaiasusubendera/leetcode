func minimizeXor(num1 int, num2 int) int {
    // count set bits ('1' in bin) of num2
    setBits2 := bits.OnesCount(uint(num2)) // setBits2: set bits of num2
    // convert int num1 into binary string
    binStr1 := strconv.FormatInt(int64(num1), 2) // binStr1: binary string of num1
    
    // copy binStr1 into slice of byte x
    binSoBX := []byte(binStr1) // binSoBX: binary slice of byte x

    // count set bits of x
    setBitsX := 0 // setBitsX: set bits of x
    for i := 0; i < len(binSoBX); i++ {
        if binSoBX[i] == '1' {
            setBitsX++
        }
    }
    
    // add extra bits '1' if needed (setBitsX < setBits2)
    n := len(binSoBX)-1
    for setBitsX < setBits2 && n >= 0 {
        if binSoBX[n] == '0' {
            binSoBX[n] = '1'
            setBitsX++
        }
        n--
    }

    // If still not enough setBits, prepend '1' to the binary string
    for setBitsX < setBits2 {
        binSoBX = append([]byte{'1'}, binSoBX...) // Add '1' to the front
        setBitsX++
    }

    // remove excess bits '1' if needed (setBitsX > setBits2)
    n = len(binSoBX)-1
    for setBitsX > setBits2 && n >= 0 {
        if binSoBX[n] == '1' {
            binSoBX[n] = '0'
            setBitsX--
        }
        n--
    }

    // convert the binary string x back to an integer
    x, _ := strconv.ParseInt(string(binSoBX), 2, 64)
    return int(x)
}

// naive solution
// time: O(log(num1) + log(num2)) or O(log(max(num1, num2)))
// space: O(log(num1))
// O(log(n)) is from bin/dec converting and set bits counting

// xor concept
// a xor a = 0, e.g. 11 xor 11 = 0 (3 xor 3 = 0)
// the more digits that are the same, the lower the xor value
// for minimal change of xor value, add/remove '1' from lsb to msb
