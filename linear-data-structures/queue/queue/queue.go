package queue

import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue() {
	var data int
	fmt.Print("enter the data to insert in queue: ")
	fmt.Scan(&data)

	q.queue = append(q.queue, data)
	fmt.Println("node inserted successfully.")
}

func (q *Queue) DisplayQueue() {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}

	fmt.Print("queue: ")
	for _, value := range q.queue {
		fmt.Printf("%d -> ", value)
	}

	fmt.Println("null")
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return -1
	}

	// example of how queue[1:] works
	// queue = [1, 2, 3, 4]
	// queue[0] = 1
	// queue[1:] = [2, 3, 4]

	// queue = [7,9,8,6]
	// queue[0] = 7
	// queue[1:] = [9,8,6]

	value := q.queue[0]   // get the first element
	q.queue = q.queue[1:] // remove the first element

	return value
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.queue[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) Size() {
	fmt.Println("size of queue: ", len(q.queue))
}

func NewQueue() Queue {
	return Queue{
		queue: []int{},
	}
}
