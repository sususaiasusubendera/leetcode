type ProductOfNumbers struct {
    prefix []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prefix: []int{1},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 { // reset the prefix product if num == 0
        this.prefix = []int{1}
    } else {
        newPre := this.prefix[len(this.prefix)-1] * num
        this.prefix = append(this.prefix, newPre)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    if k >= len(this.prefix) {
        return 0
    }
    return this.prefix[len(this.prefix)-1] / this.prefix[(len(this.prefix)-1)-k]
}

// prefix product approach
// time: O(1)
// space: O(1)


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */