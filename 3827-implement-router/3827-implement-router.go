type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	size        int
	routerQueue []Packet
	routerMap   map[Packet]struct{}

	destMap map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		size:        memoryLimit,
		routerQueue: []Packet{},
		routerMap:   make(map[Packet]struct{}),
		destMap:     make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	newPacket := Packet{source, destination, timestamp}
	if _, exist := this.routerMap[newPacket]; exist {
		return false
	}

	if this.size == len(this.routerQueue) {
		head := this.routerQueue[0]
		this.routerQueue = this.routerQueue[1:]
		delete(this.routerMap, head)

		this.removeTimestampFromDest(head.destination, head.timestamp)
	}

	this.routerQueue = append(this.routerQueue, newPacket)
	this.routerMap[newPacket] = struct{}{}
	this.destMap[destination] = append(this.destMap[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.routerQueue) == 0 {
		return []int{}
	}

	head := this.routerQueue[0]
	this.routerQueue = this.routerQueue[1:]
	delete(this.routerMap, head)
	this.removeTimestampFromDest(head.destination, head.timestamp)

	return []int{head.source, head.destination, head.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	arr := this.destMap[destination]
	if len(arr) == 0 {
		return 0
	}

	left := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= startTime
	})

	right := sort.Search(len(arr), func(i int) bool {
		return arr[i] > endTime
	})

	return right - left
}

func (this *Router) removeTimestampFromDest(destination int, timestamp int) {
	arr := this.destMap[destination]
	for i, t := range arr {
		if t == timestamp {
			this.destMap[destination] = append(arr[:i], arr[i+1:]...)
			break
		}
	}
}

// NOTICE ME SENPAI!!!
