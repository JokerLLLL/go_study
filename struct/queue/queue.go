package queue

// 重命名
//type Queue []int  int的slice

type Queue []interface{ //支持任何类型的slice
}

func (queue *Queue) Pop() interface{} {
	if queue.IsEmpty() {
		//panic("queue is empty")
		return nil
	}
	head := (*queue)[0]
	*queue = (*queue)[1:]
	//return head.(string)  拿interface里的string
	//return head.(int)     拿interface里的int
	return head
}

func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue) Push(value interface{}) {
	*queue = append(*queue, value)
}