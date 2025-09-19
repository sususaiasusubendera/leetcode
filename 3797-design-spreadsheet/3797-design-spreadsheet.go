type Spreadsheet struct {
    cellMap map[string]int
}


func Constructor(rows int) Spreadsheet {
    return Spreadsheet{ cellMap: make(map[string]int) }
}


func (this *Spreadsheet) SetCell(cell string, value int)  {
    this.cellMap[cell] = value
}


func (this *Spreadsheet) ResetCell(cell string)  {
    this.cellMap[cell] = 0
}


func (this *Spreadsheet) GetValue(formula string) int {
    f := strings.TrimPrefix(formula, "=")
    input := strings.Split(f, "+")
    a, b := input[0], input[1]

    var aInt, bInt int
    if num, ok := this.cellMap[a]; ok {
        aInt = num
    } else {
        if num, err := strconv.Atoi(a); err == nil {
            aInt = num
        } else {
            aInt = 0
        }
    }

    if num, ok := this.cellMap[b]; ok {
        bInt = num
    } else {
        if num, err := strconv.Atoi(b); err == nil {
            bInt = num
        } else {
            bInt = 0
        }
    }

    return aInt + bInt
}

// array, design, hash map, string
// time:
// - Constructor, SetCell, ResetCell, GetValue: O(1)
// space: O(n)

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */