package priority_queue

type IQueue[T any] interface {
	Enqueue(value T, priority int)
	Dequeue() (T, error)
	Peek() (T, error)
	Len() int
}
