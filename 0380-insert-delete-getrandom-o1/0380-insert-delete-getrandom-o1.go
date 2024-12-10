type RandomizedSet struct {
    Map map[int]int
    Slice []int
}


func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano())
    return RandomizedSet{
        Map: map[int]int{},
        Slice: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exist := this.Map[val]; !exist {
        this.Slice = append(this.Slice, val)
        this.Map[val] = len(this.Slice) - 1 
        return true 
    }
    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if index, exist := this.Map[val]; exist {
        lastElement := this.Slice[len(this.Slice)-1]
        this.Slice[index] = lastElement
        this.Map[lastElement] = index

        this.Slice = this.Slice[:len(this.Slice)-1]
        delete(this.Map, val)
        return true
    }
    return false
}


func (this *RandomizedSet) GetRandom() int {
    randomIndex := rand.Intn(len(this.Slice))
    return this.Slice[randomIndex]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */