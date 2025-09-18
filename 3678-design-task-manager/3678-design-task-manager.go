type TaskManager struct {
	taskMap  map[int]Task // taskId -> Task{ taskId, userId, priority }, for lazy delete
	taskHeap *TaskHeap
}

func Constructor(tasks [][]int) TaskManager {
	taskManager := TaskManager{
		taskMap:  make(map[int]Task),
		taskHeap: &TaskHeap{},
	}

	for i := 0; i < len(tasks); i++ {
		userId, taskId, priority := tasks[i][0], tasks[i][1], tasks[i][2]

		newTask := Task{
			taskId:   taskId,
			userId:   userId,
			priority: priority,
		}

		taskManager.taskMap[taskId] = newTask

		heap.Push(taskManager.taskHeap, newTask)
	}

	return taskManager
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	newTask := Task{
		taskId:   taskId,
		userId:   userId,
		priority: priority,
	}

	this.taskMap[taskId] = newTask

	heap.Push(this.taskHeap, newTask)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	userId := this.taskMap[taskId].userId
	editedTask := Task{
		taskId:   taskId,
		userId:   userId,
		priority: newPriority,
	}

	this.taskMap[taskId] = editedTask

	heap.Push(this.taskHeap, editedTask)
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskMap, taskId)
}

func (this *TaskManager) ExecTop() int {
	h := this.taskHeap // *TaskHeap
	for h.Len() > 0 { // loop until getting a valid task with valid priority
		top := (*h)[0] // curr top (Task)
		if taskFromMap, exists := this.taskMap[top.taskId]; exists { // check if the task hasn't been removed
			if top.userId == taskFromMap.userId && top.priority == taskFromMap.priority { // check if the top task is valid
                heap.Pop(h)
                this.Rmv(top.taskId)
				return top.userId
			}
		}

		heap.Pop(h) // delete top
	}

    return -1
}

// design, hash map, heap
// time:
//  - Constructor: O(nlog(n))
//  - Add, Edit, ExecTop: O(log(n))
//  - Rmv: O(1)
// space: O(n)

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

type Task struct {
	taskId   int
	userId   int
	priority int
}

// max heap implementation
type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }

func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h TaskHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].taskId > h[j].taskId
	}

	return h[i].priority > h[j].priority
}

func (h *TaskHeap) Push(x any) { *h = append(*h, x.(Task)) }

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
