type ProductOfNumbers struct {
    pre []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        pre: []int{1},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.pre = []int{1}
    } else {
        n := len(this.pre)
        this.pre = append(this.pre, this.pre[n-1]*num)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.pre)
    if k >= n {
        return 0
    }
    return this.pre[n-1] / this.pre[n-k-1]
}

// array, prefix sum
// time: O(1)
// space: O(n)

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */