func maximumScore(nums []int, k int) int {
    mod := int(1e9 + 7)
    n := len(nums)
    arr := make([][3]int, n)
    for i, x := range nums {
        arr[i] = [3]int{i, primeFactors(x), x}
    }

    left := make([]int, n)
    right := make([]int, n)
    stk := make([][2]int, 0)
    for _, fx := range arr {
        i := fx[0]
        for len(stk) > 0 && stk[len(stk)-1][0] < fx[1] {
            stk = stk[:len(stk)-1]
        }
        if len(stk) > 0 {
            left[i] = stk[len(stk)-1][1]
        } else {
            left[i] = -1
        }
        stk = append(stk, [2]int{fx[1], i})
    }

    stk = make([][2]int, 0)
    for i := len(arr) - 1; i >= 0; i-- {
        fx := arr[i]
        i := fx[0]
        for len(stk) > 0 && stk[len(stk)-1][0] <= fx[1] {
            stk = stk[:len(stk)-1]
        }
        if len(stk) > 0 {
            right[i] = stk[len(stk)-1][1]
        } else {
            right[i] = n
        }
        stk = append(stk, [2]int{fx[1], i})
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i][2] > arr[j][2]
    })

    ans := 1
    for _, fx := range arr {
        i, x := fx[0], fx[2]
        l, r := left[i], right[i]
        cnt := (i - l) * (r - i)
        if cnt <= k {
            ans = ans * pow(x, cnt, mod) % mod
            k -= cnt
        } else {
            ans = ans * pow(x, k, mod) % mod
            break
        }
    }
    return ans
}

// NOTICE ME SENPAI!!!

func pow(x, n, mod int) int {
    res := 1
    for n > 0 {
        if n&1 == 1 {
            res = res * x % mod
        }
        x = x * x % mod
        n >>= 1
    }
    return res
}

func primeFactors(n int) int {
    i := 2
    ans := make(map[int]bool)
    for i*i <= n {
        for n%i == 0 {
            ans[i] = true
            n /= i
        }
        i++
    }
    if n > 1 {
        ans[n] = true
    }
    return len(ans)
}
