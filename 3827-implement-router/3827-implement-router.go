type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	size        int
	routerQueue []Packet
	routerMap   map[Packet]struct{} // for detecting duplicate
	destMap     map[int][]int       // destination -> timestamps (increasing order)
	destIdx     map[int]int         // destination -> logical start index
}

func Constructor(memoryLimit int) Router {
	return Router{
		size:        memoryLimit,
		routerQueue: []Packet{},
		routerMap:   make(map[Packet]struct{}),
        destMap:     make(map[int][]int),
        destIdx:     make(map[int]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	newPacket := Packet{source, destination, timestamp}

	if _, exist := this.routerMap[newPacket]; exist { return false }

	if this.size == len(this.routerQueue) {
		head := this.routerQueue[0]
		this.routerQueue = this.routerQueue[1:] // dequeue
		delete(this.routerMap, head)
        this.RemoveTimestampFromDest(head.destination, head.timestamp)
	}

	this.routerQueue = append(this.routerQueue, newPacket)
	this.routerMap[newPacket] = struct{}{}
    this.destMap[destination] = append(this.destMap[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.routerQueue) == 0 { return []int{} }

	head := this.routerQueue[0]
	packet := []int{head.source, head.destination, head.timestamp}

	this.routerQueue = this.routerQueue[1:] // dequeue
	delete(this.routerMap, head)
    this.RemoveTimestampFromDest(head.destination, head.timestamp)

	return packet
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    s, ok := this.destMap[destination]
    if !ok { return 0 }

    startIdx := this.destIdx[destination]
    if startIdx >= len(s) { return 0 }

    validS := s[startIdx:]

    left := lowerBound(validS, startTime)
    right := upperBound(validS, endTime)

    return right - left
}

func (this *Router) RemoveTimestampFromDest(destination int, timestamp int) {
    s := this.destMap[destination]
    if len(s) == 0 { return }

    idx := this.destIdx[destination]
    if idx < len(s) && s[idx] == timestamp {
        this.destIdx[destination]++
    }

    newArr := this.destMap[destination]
    newIdx := this.destIdx[destination]
    if newIdx > 0 && 2*newIdx >= len(newArr) {
        this.destMap[destination] = newArr[newIdx:]
        this.destIdx[destination] = 0
    }
}

// array, binary search, design, hash map, queue
// time:
// - AddPacket: O(1)
// - ForwardPacket: O(1)
// - GetCount: O(log(n))
// - RemoveTimestampFromDest: amortized O(1)
// space: O(size + D)


/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */

// find first index >= target (binary search)
func lowerBound(s []int, target int) int {
    l, r := 0, len(s)
    for l < r {
        mid := (l + r) / 2
        if s[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

// find first index > target (binary search)
func upperBound(s []int, target int) int {
    l, r := 0, len(s)
    for l < r {
        mid := (l + r) / 2
        if s[mid] > target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}