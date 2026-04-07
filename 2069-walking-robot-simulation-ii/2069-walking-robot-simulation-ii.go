var dirs = []string{"North", "East", "South", "West"}
var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Robot struct {
	w int
	h int
	d int // 0: north, 1: east, 2: south, 3: west
	p []int
}

func Constructor(width int, height int) Robot {
	return Robot{
		w: width,
		h: height,
		d: 1,
		p: []int{0, 0},
	}
}

func (this *Robot) Step(num int) {
    cycle := 2 * (this.w + this.h) - 4
    num = num % cycle

    // edge case: num % cycle == 0 when at pos (0, 0)
    // return to (0, 0), but the dir must be handled
    if num == 0 {
        if this.p[0] == 0 && this.p[1] == 0 {
            this.d = 2 // south
        }
        return
    }

    for i := 0; i < num; i++ {
        nx, ny := this.p[0] + steps[this.d][0], this.p[1] + steps[this.d][1]
        if nx < 0 || nx >= this.w || ny < 0 || ny >= this.h {
            this.d = (this.d - 1 + 4) % 4
            nx, ny = this.p[0] + steps[this.d][0], this.p[1] + steps[this.d][1]            
        }
        this.p[0], this.p[1] = nx, ny
    }
}

func (this *Robot) GetPos() []int {
	return this.p
}

func (this *Robot) GetDir() string {
	return dirs[this.d]
}

// design, simulation
// time: O(n)
// space: O(1)

// note:
// crazy edge case

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */