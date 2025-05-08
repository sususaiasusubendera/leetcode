func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	d := make([][]int, n)
	v := make([][]bool, n)
	for i := range d {
		d[i] = make([]int, m)
		v[i] = make([]bool, m)
		for j := range d[i] {
			d[i][j] = math.MaxInt32
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	d[0][0] = 0
	q := &PriorityQueue{}
	heap.Push(q, &State{0, 0, 0})

	for q.Len() > 0 {
		s := heap.Pop(q).(*State)
		if v[s.x][s.y] {
			continue
		}
		if s.x == n-1 && s.y == m-1 {
			break
		}
		v[s.x][s.y] = true
		for _, dir := range dirs {
			nx, ny := s.x+dir[0], s.y+dir[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			dist := max(d[s.x][s.y], moveTime[nx][ny]) + (s.x+s.y)%2 + 1
			if d[nx][ny] > dist {
				d[nx][ny] = dist
				heap.Push(q, &State{nx, ny, dist})
			}
		}
	}

	return d[n-1][m-1]
}

// editorial
// NOTICE ME SENPAI!!!

type State struct {
	x, y, dis int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dis < pq[j].dis
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}