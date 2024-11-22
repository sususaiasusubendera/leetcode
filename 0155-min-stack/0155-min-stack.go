type MinStack struct {
    stack []int
    minVal []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    if len(this.minVal) == 0 || val <= this.minVal[len(this.minVal)-1] {
         this.minVal = append(this.minVal, val)
        }
    this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
    if this.minVal[len(this.minVal)-1] == this.Top() {
        this.minVal = this.minVal[:len(this.minVal)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minVal[len(this.minVal)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */