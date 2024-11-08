func twoSum(numbers []int, target int) []int {
    m := map[int]int{}
    
    for i := 0; i < len(numbers); i++ {
        c := target - numbers[i]
        
        if _, exist := m[c]; exist {
            return []int{m[c]+1, i+1}
        }

        m[numbers[i]] = i
    }
    
    return nil
}